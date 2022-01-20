This project describes an api  that accepts input as text and provides Json Output as Top ten most used words and
times of occurrence in the text.

Steps to run the application:

1. Clone the github repo for textprocessor
2. go run main.go

it will start the http server on localhost:8080

You can send post request by plaing the text in the request body to the above url:


example input text.

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed interdum pellentesque feugiat. Integer tristique, nibh ut vehicula auctor, sem lacus laoreet tellus, vel pulvinar justo orci non urna. Suspendisse potenti. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Integer iaculis elit a hendrerit finibus. Integer eget nisi varius, suscipit mi ac, egestas nibh. Fusce consectetur molestie convallis. Aliquam velit sapien, volutpat blandit massa ac, iaculis tincidunt dui. Integer at neque tellus. Nullam malesuada ante vitae tincidunt placerat. Aliquam et justo finibus, ullamcorper leo id, pharetra tortor.

output json:

[
    {
        "word": "Integer",
        "occurence": 4
    },
    {
        "word": "tincidunt",
        "occurence": 2
    },
    {
        "word": "ac,",
        "occurence": 2
    },
    {
        "word": "justo",
        "occurence": 2
    },
    {
        "word": "ante",
        "occurence": 2
    },
    {
        "word": "Aliquam",
        "occurence": 2
    },
    {
        "word": "orci",
        "occurence": 2
    },
    {
        "word": "ipsum",
        "occurence": 2
    },
    {
        "word": "consectetur",
        "occurence": 2
    },
    {
        "word": "et",
        "occurence": 2
    }
]
