Here to figure out what the Golang structure of this application is going to be.
Influenced by the [Writing Web Applications](golang.org/doc/articles/wiki/) 
example found in the [official documentation](golang.org).

The guide starts off by explaining the goal.

What is it that we are trying to build, the goal?

    "A wiki consists of a series of interconnected pages, each of whch has a
    title and a body (the page content). Here, we define Page as a struct with
    two fields representing the title and body."

What is OUR goal?

    Budget is a fancy word for a record of transactions, with associated
    metadata, e.g. categories, card, etc. As a foundation, we will start with
    3 basic components of a budget: transactions, categories, and cards.

    Transaction - a struct with transaction metadata
        ID
        Name
        AmountUSD
        CardID
        CategoryID

    Card - struct with card metadata
        ID
        Name
        Owner

    Category - struct with category metadata
        ID
        Name
        IsEssential

    A key assumption here is your program will be recieving these structs
    through an API call. (subject to change) Our goal is to setup the API
    call.


# Design

How am I going to be recieving the data from the front-end?

[Front-end](<insert_repo_link>)

 
 This is done through the front-end sending a POST request + json payload to one
 of the API endpoints. In attempt to accomplish exactly this, I'm going to make
 a page that has a botton. When clicked, a POST request will be sent.
