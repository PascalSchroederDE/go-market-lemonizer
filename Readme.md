# Market Lemonizer

This library aims towards simplifying the access to the [lemon market api](https://www.lemon.markets) for Golang developers. It provides simple interfaces for basic operations/interactions with the lemon market api and returns golang structs based on which further calculations/operations can be executed. It's free to use for everybody and contributions are highly welcome.

## General

### Warnings

**FUNCTIONALITY WARNING**: As this is a private project I cannot guarantee it's functionality at all time and for all purposes. Also I have not included all interfaces and all parameters (yet). This might follow in a future version

**API CALL WARNING**: As this library aims to simplify the usage of the lemon market api as good as possible it also tries to simplify every single operation for a good and very easy usage. This results in the possibility, that for one data request many API calls will be executed, which might have an influence on your lemon market api (usage) plan. To be exact: There are no limitations in number of ISINs, number of returned data points or time spans when requesting data via this api. On the other side these very parameters are limited per call towards the lemon market api, e.g. you can only request 10 ISINs at a time and get at most 100 data points returned per page (for each page you need a seperate call). Also you can only request data points per minute or hour (h1 or m1) within the same day with this API. The library simplifies this by having no limitations on the parameters and executing as many calls as needed for the requested options, which may result in more API calls than expected. To give an example: When you want to get all data points per hour for 17 ISINs for the last two days, the ISINs will be splitted into **2** sets. We have 120 data points per day for the first ISIN package (meaning 2 pages) and 70 for the second (meaning 1 page) for both days. This will result in 2 (pages for first ISIN set per day) * 2 (days) + 1 (pages for the second ISIN set per day) * 2 (days) = 6 API calls for this one request.

### Why should I use this library?

As already described this library aims on simplifying Lemon market API calls as much as possible. As the API itself has some (necessary) limitations per API call (as described in Warnings/API Call Warning) this library simplifies them to work around those limitations and remove the hurdle of implementing every single API call by yourself for gaining all the data you need.

Also this library comes with Golang structs for each response to have the data prepared for further calculatiosn and operations. Those structs will be explained later in this documentation.

In the end it is a very easy to use library and fulfills many needs to simplify your application.

### Limitations

On the other side this library is not finished yet and comes with some limitations to the functionalities of the API. 

For now not every endpoint which is provided by the Lemon Market API has a translation/implementation within this library and only the in my view most important ones are implemented.

Also only the most important parameters have been considered and not every possible parameter to be included in the request can also be included in a request via this library.

Last please read the Warnign regarding API Call limitation to be aware of the number of API calls one request executed via this library might cause. 

### Structure

This library basically consists out of 3 main available packages, which are

- account-info
- market-data
- trading

Those 3 packages provide access to different parts of the lemon market api, whereby account-info includes information about your personal account and your portfolio. Market-data provides historic (and soon realtime) data of the market for all the available titles, either per minute, per hour or per day, and returns OHLC (Opening/High/Low/Closing) information about the prices for all available data points. Last trading enables you to place buy or sell orders on your lemon api account and activate or delete them.

## How to

### How to sign up for lemon market api

For signing up to lemon market api please have a look at their web page at [lemon.markets](https://www.lemon.markets).

### How to install

For installing this library please execute following command within your go project

```
go install github.com/PascalSchroederDE/go-market-lemonizer
```

Then you need to import the necessary parts of this library. For being able to handle the returned go structures you need to import following:

```
lemonStructs "github.com/PascalSchroederDE/go-market-lemonizer/pkg/lemonizer/datastructs"
lemonInterval "github.com/PascalSchroederDE/go-market-lemonizer/pkg/lemonizer/datastructs/interval"
```

While lemonStructs is used in general for all returned responses, lemonInterval is especially dedicated to the OHLC history data requests. With this import you can differentiate whether you would like to get data per minute, hour or day. For further description please look up in chapter *How to use*.

To then use the different packages (account-info/market-data/trading) of this library, you can import them as following:

```
lemonData "github.com/PascalSchroederDE/go-market-lemonizer/pkg/lemonizer/api/market-data"
lemonAccount "github.com/PascalSchroederDE/go-market-lemonizer/pkg/lemonizer/api/account-info"
lemonTrading "github.com/PascalSchroederDE/go-market-lemonizer/pkg/lemonizer/api/trading"
```

From here on you can continue to initialize accordig structures to being able to communicate with the lemon market API.

### How to initialze

For initializing a working communication point from your golang project to the lemon market API via this library you can use existing interfaces to create an according go-struct. For that we assume, that we have imported our packages as described in *How to install*. Then each package provide following functions:

- Init(apikey string): Initialize with default lemon market endpoint (given as const in config of this library) by only your apikey necessary to be passed
- InitManualEndpoint(apikey string, endpoint string): Initialize with a manually specified endpoint needed to be passed additionally to the apikey. Should be a complete url starting with `https://` and should only consist out of the root path.

The lemonAccount and lemonTrading packages additionally provide following function:

- InitPaperEndpoint(apikey string): Initialize with default lemon market endpoint for paper trading (given as const in config of this library) by only your apkey necessary to be passed.

Anyway you will need your apikey, for how to get this please [lookup in lemon market api documentation](https://dashboard.lemon.markets/signup). 

Initializing those packages can then look as following:

```
apiKey := os.Getenv("LEMON_API_KEY")
if apiKey == "" {
    logger.Error("Lemon API Key us not given. Exiting")
    os.Exit(2)
}

lemonDataStruct := lemonData.Init(apiKey)

//For using paper account data, to use real account data please use `Init` as with lemonDataStruct
lemonAccountStruct := lemonAccount.InitPaperEndpoint(apiKey)

//In case you have some local trading mockup (otherwise please initialize the same way as lemonAccountStruct (for paper trading) or lemonDataStruct (for real trading))
apiBaseUrl := https://paper-trading.localhost/
lemonTradingStruct := lemonTrading.InitManualEndpoint(apiKey, apiBaseUrl)
```

Please note that usually `Init` or `InitPaperEndpoint` should be sufficient also for lemonTrading, I just added the third option for some local mockups or in case the endpoint will change and this library is not updated with the new endpoint yet. Please dont use this as copy and paste for your project, but rather use either `Init` or `InitPaperEndpoint` for all packages.

### How to use

### How to handle responses

## Example usages

### Account Info

### Market Data

### Trading

## Contribute

### Pull-Request-Workflow

### Future Plans

