package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type pageInfo struct {
	StatusCode int
	Links      map[string]int
}

func handler(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.Query().Get("url")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	log.Println("visiting", URL)

	c := colly.NewCollector()

	p := &pageInfo{Links: make(map[string]int)}

	// count links
	// c.OnHTML(".book-list-wrapper", func(e *colly.HTMLElement) {
	// 	link := e.ChildAttr("a", "title") + "**" + e.Request.AbsoluteURL(e.ChildAttr("a", "href")) + "**" + e.ChildText("p") + "**" + e.ChildAttr("img", "src")
	/*Rokomari Start */
	c.OnHTML(`.col-xl-3`, func(e *colly.HTMLElement) {
		link :=
			e.ChildText(".book-title") + "**" +
				e.ChildText(".book-author") + "**" +
				e.ChildText(".book-status") + "**" +
				e.ChildText(".book-price>strike") + "**" +
				e.ChildText(".book-price>span") + "**" +
				e.Request.AbsoluteURL(e.ChildAttr("a", "href")) + "**" +
				e.ChildAttr("img", "data-src")

		// title := e.ChildAttr("a", "title")
		// price := e.ChildText("p")
		// img := "https" + e.ChildAttr("img", "src")
		if link != "" {
			p.Links[link]++
		}
		// if title != "" {
		// 	p.Links[title]++
		// }
		// if price != "" {
		// 	p.Links[price]++
		// }
		// if img != "" {
		// 	p.Links[img]++
		// }

		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://saad:1234@cluster0.1uqmq.mongodb.net/<dbname>?retryWrites=true&w=majority"))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)

		defer client.Disconnect(ctx)

		quickstartDatabase := client.Database("rokomari")
		booksCollection := quickstartDatabase.Collection("books")
		// episodesCollection := quickstartDatabase.Collection("episodes")

		booksCollection.InsertOne(ctx, bson.D{
			{"all", link},
		})

		// episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
		// 	bson.D{
		// 		{"podcast", podcastResult.InsertedID},
		// 		{"title", "GraphQL for API Development"},
		// 		{"description", "Learn about GraphQL from the co-creator of GraphQL, Lee Byron."},
		// 		{"duration", 25},
		// 	},
		// 	bson.D{
		// 		{"podcast", podcastResult.InsertedID},
		// 		{"title", "Progressive Web Application Development"},
		// 		{"description", "Learn about PWA development with Tara Manicsic."},
		// 		{"duration", 32},
		// 	},
		// })
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Inserted %v documents into episode collection!\n", len(booksResult.InsertedIDs))

		println(link)

	})
	/*Rokomari END */

	/*BD News 24 Start*/

	c.OnHTML(`.category-area`, func(e *colly.HTMLElement) {
		link :=
			e.ChildText(".list>a") + "**" +
				e.ChildText(".list>p") + "**" +
				e.ChildText("time") + "**" +
				e.Request.AbsoluteURL(e.ChildAttr("a", "href")) + "**" +
				"https://www.banglanews24.com/" + e.ChildAttr("img", "data-src")

		if link != "" {
			p.Links[link]++
		}
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://saad:1234@cluster0.1uqmq.mongodb.net/<dbname>?retryWrites=true&w=majority"))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)

		defer client.Disconnect(ctx)

		quickstartDatabase := client.Database("rokomari")
		booksCollection := quickstartDatabase.Collection("news")

		booksCollection.InsertOne(ctx, bson.D{
			{"all", link},
		})
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Inserted %v documents into episode collection!\n", len(booksResult.InsertedIDs))

		println(link)

	})
	/* BD NEWS END */

	/*Gadget SHop BD  Start*/

	c.OnHTML(`.product-grid`, func(e *colly.HTMLElement) {
		link :=
			e.ChildText(".name>a") + "**" +
				e.ChildText(".description") + "**" +
				e.ChildText(".price") + "**" +
				e.Request.AbsoluteURL(e.ChildAttr("a", "href")) + "**" +
				e.ChildAttr("img", "src")

		if link != "" {
			p.Links[link]++
		}
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://saad:1234@cluster0.1uqmq.mongodb.net/<dbname>?retryWrites=true&w=majority"))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)

		defer client.Disconnect(ctx)

		quickstartDatabase := client.Database("rokomari")
		booksCollection := quickstartDatabase.Collection("gadgets")

		booksCollection.InsertOne(ctx, bson.D{
			{"all", link},
		})
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Inserted %v documents into episode collection!\n", len(booksResult.InsertedIDs))

		println(link)

	})
	/* Gadget SHop BD  END */

	/*Book SHop BD  Start*/

	c.OnHTML(`.type-product`, func(e *colly.HTMLElement) {
		link :=
			e.ChildText(".name>a") + "**" +
				// e.ChildText(".description") + "**" +
				e.ChildText(".price") + "**" +
				e.Request.AbsoluteURL(e.ChildAttr("a", "href")) + "**" +
				e.ChildAttr("img", "data-src")

		if link != "" {
			p.Links[link]++
		}
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://saad:1234@cluster0.1uqmq.mongodb.net/<dbname>?retryWrites=true&w=majority"))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)

		quickstartDatabase := client.Database("rokomari")
		booksCollection := quickstartDatabase.Collection("books2")
		//url: https://bookshopbd.com/c/engineering-textbooks/

		booksCollection.InsertOne(ctx, bson.D{
			{"all", link},
		})
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Inserted %v documents into episode collection!\n", len(booksResult.InsertedIDs))

		println(link)

	})
	/* bookshop BD  END */

	/*Boi Bazar BD  Start*/

	c.OnHTML(`.tab-pane`, func(e *colly.HTMLElement) {
		link :=
			e.ChildText(".line_nowrap_prod >p") + "**" +
				e.ChildText(".line_nowrap_aut>div>p") + "**" +
				e.ChildText(".tk_home") + "**" +
				e.Request.AbsoluteURL(e.ChildAttr("a", "href")) + "**" +
				e.ChildAttr(".fixed-container-book>img", "lsrc")

		if link != "" {
			p.Links[link]++
		}
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://saad:1234@cluster0.1uqmq.mongodb.net/<dbname>?retryWrites=true&w=majority"))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)

		quickstartDatabase := client.Database("rokomari")
		booksCollection := quickstartDatabase.Collection("books3")
		//url: https://www.boibazar.com/category-books/mathematics-science-technology-001

		booksCollection.InsertOne(ctx, bson.D{
			{"all", link},
		})
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Inserted %v documents into episode collection!\n", len(booksResult.InsertedIDs))

		println(link)

	})
	/* Boi Bazar BD  END */

	// extract status code
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
		p.StatusCode = r.StatusCode
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
		p.StatusCode = r.StatusCode
	})

	c.Visit(URL)

	// dump results
	b, err := json.Marshal(p)
	if err != nil {
		log.Println("failed to serialize response:", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	// example usage: curl -s 'http://127.0.0.1:7171/?url=http://go-colly.org/'
	// example usage: curl -s 'http://127.0.0.1:7171/?url=http://go-colly.org/'
	//http://127.0.0.1:7171/?url=https://cou.ac.bd/cse/facultymember

	//http://127.0.0.1:7171/?url=https://rokomari.com/books
	//https://www.banglanews24.com/category/%E0%A6%9C%E0%A6%BE%E0%A6%A4%E0%A7%80%E0%A7%9F/1
	//http://gadgetshopbd.com/Gadgets
	addr := ":7171"

	http.HandleFunc("/", handler)

	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
