package rest

import (
	"context"
	"net/http"
	"testing"
)

// 5.95KB example file
var jsonSample = `
[
  {
    "_id": "5f60877bcf9be927093ab0f8",
    "index": 0,
    "guid": "a15a6c11-d179-4dba-8f7f-1201b8f78d2c",
    "isActive": false,
    "balance": "$1,594.41",
    "picture": "http://placehold.it/32x32",
    "age": 28,
    "eyeColor": "green",
    "name": "Hoover Browning",
    "gender": "male",
    "company": "AQUAFIRE",
    "email": "hooverbrowning@aquafire.com",
    "phone": "+1 (848) 500-2512",
    "address": "772 Stewart Street, Gorham, Nevada, 4507",
    "about": "Nostrud reprehenderit irure laboris consectetur fugiat nostrud cillum et fugiat consequat aliqua. Aute veniam quis consectetur aute exercitation magna duis. Nisi id nisi cillum ullamco mollit. Velit fugiat et aliquip ea laboris sit elit velit aute reprehenderit.\r\n",
    "registered": "2017-01-23T12:16:30 -07:00",
    "latitude": 37.515602,
    "longitude": 63.877497,
    "tags": [
      "dolore",
      "consectetur",
      "ex",
      "consectetur",
      "aliqua",
      "id",
      "Lorem"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Woodard Gilliam"
      },
      {
        "id": 1,
        "name": "Reed Torres"
      },
      {
        "id": 2,
        "name": "Horn Sampson"
      }
    ],
    "greeting": "Hello, Hoover Browning! You have 2 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5f60877b6a941012166c2cae",
    "index": 1,
    "guid": "31544b0a-a2e0-4e99-9cdf-5ab1df06703d",
    "isActive": false,
    "balance": "$1,371.27",
    "picture": "http://placehold.it/32x32",
    "age": 22,
    "eyeColor": "green",
    "name": "Megan May",
    "gender": "female",
    "company": "ZAPPIX",
    "email": "meganmay@zappix.com",
    "phone": "+1 (980) 567-3042",
    "address": "966 Coleman Street, Yogaville, West Virginia, 2007",
    "about": "Commodo anim cillum elit in. Exercitation sint officia dolor dolor incididunt do deserunt dolore adipisicing. Culpa Lorem sit consequat duis. Aliqua dolor exercitation esse nulla ad. Ipsum tempor qui incididunt esse incididunt.\r\n",
    "registered": "2020-02-20T11:39:14 -07:00",
    "latitude": 85.848952,
    "longitude": -30.672706,
    "tags": [
      "irure",
      "sint",
      "commodo",
      "consectetur",
      "sint",
      "adipisicing",
      "nulla"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Susanna Orr"
      },
      {
        "id": 1,
        "name": "Kerry Bass"
      },
      {
        "id": 2,
        "name": "Erin Nichols"
      }
    ],
    "greeting": "Hello, Megan May! You have 3 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "5f60877b314209a41c566474",
    "index": 2,
    "guid": "b9ef864b-057b-426f-9b10-0c334df706bf",
    "isActive": false,
    "balance": "$2,387.88",
    "picture": "http://placehold.it/32x32",
    "age": 24,
    "eyeColor": "blue",
    "name": "Jennie Strickland",
    "gender": "female",
    "company": "MAXIMIND",
    "email": "jenniestrickland@maximind.com",
    "phone": "+1 (984) 555-3883",
    "address": "972 Tennis Court, Elwood, Colorado, 9052",
    "about": "Magna ex nostrud occaecat tempor esse duis consequat laboris. Esse fugiat quis mollit enim. Officia sunt qui veniam tempor esse eu. Esse officia occaecat laborum velit culpa pariatur mollit laborum magna aliqua.\r\n",
    "registered": "2019-05-29T03:13:18 -07:00",
    "latitude": 21.061151,
    "longitude": -92.455606,
    "tags": [
      "ut",
      "aliquip",
      "voluptate",
      "pariatur",
      "nostrud",
      "pariatur",
      "ex"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Yang Swanson"
      },
      {
        "id": 1,
        "name": "Potts Coffey"
      },
      {
        "id": 2,
        "name": "Nikki Norton"
      }
    ],
    "greeting": "Hello, Jennie Strickland! You have 4 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5f60877bfed3d709207a1e61",
    "index": 3,
    "guid": "97f31df7-4fdb-4782-b6c0-17144b7adbe8",
    "isActive": false,
    "balance": "$1,497.26",
    "picture": "http://placehold.it/32x32",
    "age": 31,
    "eyeColor": "brown",
    "name": "Cindy Booker",
    "gender": "female",
    "company": "XYQAG",
    "email": "cindybooker@xyqag.com",
    "phone": "+1 (923) 418-3967",
    "address": "796 Seigel Court, Williston, Nebraska, 8107",
    "about": "Cillum duis adipisicing est mollit aliqua exercitation. Commodo et voluptate quis pariatur ad culpa anim. Mollit enim do ullamco nisi pariatur commodo ea cillum sint elit.\r\n",
    "registered": "2017-09-16T06:07:00 -07:00",
    "latitude": -0.637079,
    "longitude": 126.451298,
    "tags": [
      "elit",
      "duis",
      "Lorem",
      "veniam",
      "ullamco",
      "nisi",
      "anim"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Carlene Patton"
      },
      {
        "id": 1,
        "name": "Earlene Avila"
      },
      {
        "id": 2,
        "name": "Dorothy Mcintosh"
      }
    ],
    "greeting": "Hello, Cindy Booker! You have 8 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5f60877bea7036d117c28c1b",
    "index": 4,
    "guid": "f9a95cf9-9001-4ad6-a456-dcad72655210",
    "isActive": true,
    "balance": "$3,005.00",
    "picture": "http://placehold.it/32x32",
    "age": 20,
    "eyeColor": "brown",
    "name": "Wheeler Reeves",
    "gender": "male",
    "company": "COMTOUR",
    "email": "wheelerreeves@comtour.com",
    "phone": "+1 (823) 542-3145",
    "address": "737 Jamaica Avenue, Fillmore, Guam, 5585",
    "about": "Velit anim nisi irure id nisi amet elit deserunt excepteur excepteur ad. Commodo veniam consequat sit fugiat. Minim pariatur consectetur cillum nostrud sunt ut proident amet minim qui et. Nulla est exercitation dolor voluptate veniam. Voluptate non aliquip commodo non sunt esse et ea qui sunt ad. Et anim in consequat sint proident eu sunt officia. Tempor voluptate pariatur velit in.\r\n",
    "registered": "2017-10-13T11:28:06 -07:00",
    "latitude": 63.160928,
    "longitude": 92.429022,
    "tags": [
      "ea",
      "reprehenderit",
      "magna",
      "consequat",
      "enim",
      "et",
      "ipsum"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Francisca Pate"
      },
      {
        "id": 1,
        "name": "Compton Grant"
      },
      {
        "id": 2,
        "name": "Marcella Hyde"
      }
    ],
    "greeting": "Hello, Wheeler Reeves! You have 1 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5f60877bc46000a0e8f20a2e",
    "index": 5,
    "guid": "cddbe724-5dcd-4dbc-afa4-7b10a6fb5ac1",
    "isActive": true,
    "balance": "$2,995.68",
    "picture": "http://placehold.it/32x32",
    "age": 39,
    "eyeColor": "green",
    "name": "Williams Long",
    "gender": "male",
    "company": "STREZZO",
    "email": "williamslong@strezzo.com",
    "phone": "+1 (890) 474-3337",
    "address": "595 Schenck Avenue, Witmer, Vermont, 9221",
    "about": "Amet anim anim est adipisicing sint irure incididunt Lorem aute. Dolor dolore ut aute exercitation minim Lorem cupidatat deserunt dolore sunt. Eu exercitation qui et qui fugiat eu quis excepteur. Quis elit sint amet adipisicing tempor officia dolore officia et sint. Aliquip consequat esse ullamco sint ad. Velit velit ipsum non cillum eiusmod minim duis est aute eiusmod ut ut mollit ea.\r\n",
    "registered": "2017-03-10T04:28:24 -07:00",
    "latitude": 8.429012,
    "longitude": -95.882877,
    "tags": [
      "dolore",
      "velit",
      "enim",
      "non",
      "dolor",
      "elit",
      "ex"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Park Faulkner"
      },
      {
        "id": 1,
        "name": "Bradley Kane"
      },
      {
        "id": 2,
        "name": "Mullen Moses"
      }
    ],
    "greeting": "Hello, Williams Long! You have 10 unread messages.",
    "favoriteFruit": "banana"
  }
]
`

// BenchmarkDoHttpCall see the performance of HTTP Call POST using this lib without no hooks
func BenchmarkDoHttpCall(b *testing.B) {
	testClient := &mockClient{
		DoFunc: doFuncMock([]byte(`{}`), nil),
	}

	client, err := DefaultClient(testClient)
	if err != nil {
		b.Error(err)
		b.FailNow()
		return
	}

	for i := 0; i <= b.N; i++ {
		_, err := client.Post(
			context.Background(),
			"", "http://example.com/",
			http.Header{},
			nil,
		)

		if err != nil {
			b.Error(err)
			b.FailNow()
			return
		}
	}

}

// BenchmarkMultipleReadResponseBody see the performance of response read.
// Since this code can unmarshal the response body multiple times, we want to see the performance effect.
func BenchmarkMultipleReadResponseBody(b *testing.B) {
	ctx := context.Background()

	testClient := &mockClient{
		DoFunc: doFuncMock([]byte(`{}`), nil),
	}

	client, err := DefaultClient(testClient)
	if err != nil {
		b.Error(err)
		b.FailNow()
		return
	}

	resp, err := client.Post(
		ctx,
		"", "http://example.com/",
		http.Header{},
		nil,
	)

	if err != nil {
		b.Error(err)
		b.FailNow()
		return
	}

	var out interface{}

	for i := 0; i <= b.N; i++ {
		err = resp.ToJson(ctx, &out)
		if err != nil {
			b.Error(err)
			b.FailNow()
			return
		}
	}

}
