package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/xellio/whois"
	"golang.org/x/net/html"

	goose "github.com/advancedlogic/GoOse"
)

// LogicModel is
type LogicModel struct {
}

/*GetInformationFromServers is a method that: search the server information
and save this information in the DB
*/
func (logicModel LogicModel) GetInformationFromServers(domainToSearch string) Domain {

	db, err := sql.Open("postgres", "postgresql://apirestuser@localhost:26257/apirest?sslmode=disable")

	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	var sslGradeDomain string
	urlServerInformation := "https://api.ssllabs.com/api/v3/analyze?host=" + domainToSearch

	fmt.Println("url: " + urlServerInformation)

	res, err := http.Get(urlServerInformation)

	if err != nil {
		fmt.Print(err)
	}

	defer res.Body.Close()

	domainObject := Domain{}
	resp, err := ioutil.ReadAll(res.Body)

	_ = json.Unmarshal(resp, &domainObject)

	for i := 0; i < len(domainObject.Servers); i++ {
		sslGradeDomain = domainObject.Servers[i].SslGrade
		if strings.Compare(sslGradeDomain, domainObject.Servers[i].SslGrade) == -1 {
			sslGradeDomain = domainObject.Servers[i].SslGrade
		}
	}

	domainObject.SslGrade = sslGradeDomain

	// Insert information to the DataBase
	for i := 0; i < len(domainObject.Servers); i++ {
		fmt.Println("saving" + domainObject.Servers[i].Address)
		query := "INSERT INTO apirest.domainservers (domain, servers, sslgrade, timeAccess) VALUES ('" + domainToSearch + "', '" + domainObject.Servers[i].Address + "', '" + domainObject.SslGrade + "', '" + strconv.FormatInt(time.Now().Unix(), 10) + "')"
		if _, err := db.Exec(
			query); err != nil {
			fmt.Println("Error")
			fmt.Println(err)
		}
	}

	// Get information about the servers, OrgName, and Country
	for i := 0; i < len(domainObject.Servers); i++ {
		ip := net.ParseIP(domainObject.Servers[i].Address)
		res, err := whois.QueryIP(ip)
		if err != nil {
			fmt.Println(err)
			return Domain{}
		}

		owner := strings.Join(res.Output["OrgName"], ",")
		country := strings.Join(res.Output["Country"], ",")

		domainObject.Servers[i].Country = country
		domainObject.Servers[i].Owner = owner
	}

	rows, err := db.Query("Select * from apirest.domainservers ORDER BY timeaccess")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	domainObject.ServerChanged = false

	for rows.Next() {

		var domain, servers, sslgrade string
		var timeaccess int64
		if err := rows.Scan(&domain, &servers, &sslgrade, &timeaccess); err != nil {
			log.Fatal(err)
		}

		if strings.Compare(domainToSearch, domain) == 0 {
			if (time.Now().Unix() - timeaccess) >= 3600 {
				domainObject.PreviousSSLGrade = sslgrade
			} else {
				domainObject.PreviousSSLGrade = domainObject.SslGrade
			}
			domainObject.ServerChanged = !inSliceServer(domainObject.Servers, servers)
		}
	}
	return domainObject
}

/*GetInformationFromDomain is a method that look for the Title and Logo
of the domain from the html
*/
func (logicModel LogicModel) GetInformationFromDomain(domain Domain, domainName string) Domain {

	url := "https://www." + domainName

	res, err := http.Get(url)

	if err != nil {
		fmt.Print(err)
		domain.IsDown = true
		return domain
	}

	defer res.Body.Close()

	htmlPlainText, err := ioutil.ReadAll(res.Body)

	r := strings.NewReader(string(htmlPlainText))

	doc, err := html.Parse(r)

	if err != nil {
		panic("fail to parse")
	}

	title, _ := traverseTree(doc)

	domain.Title = title

	g := goose.New()
	art, _ := g.ExtractFromURL(url)

	logo := art.TopImage

	domain.Logo = logo

	return domain
}

/* This method look for the tag title in an html node */
func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

/* This method is in charge of move trough the tree */
func traverseTree(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		result, ok := traverseTree(i)
		if ok {
			return result, ok
		}
	}

	return "", false
}

// GetListedServers is in charge of list all the servers a client has search previously
func (logicModel LogicModel) GetListedServers() Item {

	db, err := sql.Open("postgres", "postgresql://apirestuser@localhost:26257/apirest?sslmode=disable")

	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	item := Item{}

	rows, err := db.Query("Select domain from apirest.domainservers")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var domain string
		if err := rows.Scan(&domain); err != nil {
			log.Fatal(err)
		}

		if !inSlice(item.Domains, domain) {
			item.Domains = append(item.Domains, domain)
		}
	}

	for i := 0; i < len(item.Domains); i++ {
		fmt.Println(item.Domains[i])
	}
	return item
}

/* This method returns if a value is inside an array */
func inSlice(arr []string, val string) bool {
	for _, v := range arr {
		if strings.Compare(v, val) == 0 {
			return true
		}
	}
	return false
}

/* This method returns if a value is inside an array */
func inSliceServer(arr []Server, val string) bool {
	for _, v := range arr {
		if strings.Compare(v.Address, val) == 0 {
			return true
		}
	}
	return false
}
