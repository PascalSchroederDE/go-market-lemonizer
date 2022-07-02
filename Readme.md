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

### How to use / Function Overview

After initializing the library can be easily used by calling the appropriate methods with belongig parameters. Following methods are available:

- lemonData
    - GetLatestData(isin []string, interval lemon_interval.LemonInterval): map[string][]datastructs.OHLCData
        - Requesting latest data by giving a (arbitrary sized) list of ISINs as string  and the lemonInterval (Minutes, Hours or Days).
        - Returns the latest data/values of given interal for all ISINs as OHLCData object (includes Opening, High, Low & Closing prices)
    - GetHistoryData(isin []string, dateFrom time.Time, dateTo time.Time, interval lemon_interval.LemonInterval: map[string][]datastructs.OHLCData
        - Requesting history data by giving a (arbitrary sized) list of ISINs as string, a beginning and ending timestamp  and the lemonInterval (Minutes, Hours or Days)
        - Returns the latest data/values of the given timespan in the given interal for all ISINs as OHLCData object (includes Opening, High, Low & Closing prices)
    - GetVenueInformation(mic string): datastructs.VenueResult
        - Requesting informations of a venue by giving the MIC as string
        - Returns venue information such as OpeningHours, IsOpen Boolean and Venue-Title
- lemonTrading
    - PlaceSellStockTitleOrder(isin string, amount int, limitPrice int, expiryTime time.Time): datastructs.PlaceOrderResults
        - Requesting to place a sell order by giving the isin, the amount of titles to be sold, the limit price for the order and the expiry time when the order is not valid anymore
        - Returns a PlaceOrderResults object which includes a summary of the placed order and also the id, which is needed to activate, delete or track the order
    - PlaceBuyStockTitleOrder(isin string, amount int, limitPrice int, expiryTime time.Time): datastructs.PlaceOrderResults
        - Requesting to place a buy order by giving the isin, the amount of titles to be sold, the limit price for the order and the expiry time when the order is not valid anymore
        - Returns a PlaceOrderResults object which includes a summary of the placed order and also the id, which is needed to activate, delete or track the order
    - ActivateOrder(orderId string)
        - Requesting to active an order by giving the id, which was included in the PlaceOrderResults object returned by the Place[...]StockTitleOrder function. **WARNING** Once an order is being activated costs might be charged
        - Returns an error if activation fails. If no error is returned it was successful
    - DeleteOrder(orderId string)
        - Requesting to delete an order by giving the id, which was included in the PlaceOrderResults object returned by the Place[...]StockTitleOrder function
        - Returns an error if deletion fails. If no error is returned it was successful
- lemonAccount
    - GetPortfolio(): (map[string]datastructs.PortfolioResults
        - Requesting currently possessed titles in your account. Returns map which allocated the ISIN of the title to its Portfolio information like ISIN-Title, Quantities and Buy Prices
    - GetAccount(): datastructs.GetAccountResult
        - Requesting account information such as creation date, name but also balance or trading plan

All functions will return either the described response object or an error if the request has failed. This can be checked as following:

```
responseObj, err := lemonizerAPIObj.ArbitraryFunction(parameters...)
if err != nil {
    logger.Error(err)
    os.Exit(2)
}
```

### How to handle responses / Response objects

All above described functions return either an error or most return a datastruct with valuable information about the request results. They are all included in the package

```
lemonStructs "github.com/PascalSchroederDE/go-market-lemonizer/pkg/lemonizer/datastructs"
```

To get a short overview of which datastructs exist see following list:

- Account
    - PortfolioResults: Includes information about the currently possessed stock titles such as Buy prices, amount, ISIN and title
    - GetAccountResult: Includes information about the account such as Owner name, E-Mail adress, creation date, billing information, balance, trading plan or tax information
- Data
    - OHLCData: Includes all OHLC-data (Opening, Closing, Highest and Lowest price/value) of a specific title as well a timestamp of these OHLC information
    - VenueResult: Includes information of a requested venue such as Name, mic or opening information
    - OpeningHours: Included in the VenueResults datastruct. Includes starting time, closing time and timezone information
- Trading
    - PlaceOrderResults: Includes summary of placed order such as OrderID, status, quantity, expiry time, venue, price information or regulatory information
    - RegulatoryInformation: Included in the PlaceOrderResults datastruct. Includes regulatory information of the placed order such as Costs-Entry or legal disclaimer

For more detailed information please have a direct look inside of the `/pkg/datastructs` directory.

To access single fields of the datastructs you can use default golang syntax, e.g.

```
func getOpeningValue(datapoint lemonStructs.OHLCData) datastructs.Datapoint {
	return datapoint.Opening
}
```

## Example usages

This section contains some example code snippets of how this library can be used. For information about how to initialize the API in the first place please look in above section *How to initialze*

### Account Info

```
func GetBuyingPower() (int, error) {
	marketLemonizerAPI := initAccountAPIStruct()

	spaceRespObj, err := marketLemonizerAPI.GetAccount()
	if err != nil {
		return 0, errors.New("account could not be grapped")
	}
	return spaceRespObj.CashToInvest, nil
}

func IsVenueOpen(mic string) bool {
	marketLemonizerAPI := initAPIStruct()

	venueInformation, err := marketLemonizerAPI.GetVenueInformation(mic)
	if err != nil {
		return false
	}

	return venueInformation.IsOpen
}
```

### Market Data

```
func getLastDayChartHistories(isins []string) map[string][]lemonDatastructs.OHLCData, error {
	marketLemonizerAPI := initAPIStruct()

	dateFrom := baseTime.Add(-24 * time.Hour)
	dateTo := baseTime

	requestedStockData, err := marketLemonizerAPI.GetHistoryData(isins, dateFrom, dateTo, lemonInterval.Hour)
	if err != nil {
		return make(map[string][]lemonDatastructs.OHLCData), errors.New("no data")
	}

	return transformedStockData, nil
}
```

### Trading

```
order := make(map[string]projectDatastructs.OrderInfos)
order["US88160R1014"]:=projectDatastructs.OrderInfos{LimitPrice: 10, Amount: 1, ExpiryTime: time.Now().Add(1*time.Hour)}
for isin, orderInformation := range orders {
    limitPrice := int(orderInformation.LimitPrice)
    placedOrderObj, err := marketLemonizerAPI.PlaceBuyStockTitleOrder(isin, orderInformation.Amount, limitPrice, orderInformation.ExpiryTime)
    if err != nil {
        logger.Warn("Buy order not successfully created: " + isin)
    } else {
        logger.Info("Buy order successfully created: " + isin)
        err := marketLemonizerAPI.ActivateOrder(orderId)
        if err != nil {
            logger.Info("Activated order")
        } else {
            marketLemonizerAPI.DeleteOrder(orderId)
        }
    }
}
```

## Contribute

### Pull-Request-Workflow

You are more than welcome to contribute to this project with your own code. New features or bugfixes are always welcome. For this purpose please create a new branch with following name specifications:

<type-of-contribution>/<short-description>

As type of contribution please use one of following which fits best in your opinion:

- bugfix
- feature
- documentation
- api_update

As short description please summarize in a few words roughly what your contribution is about.

Once your code changes are done please push this branch and request a Review by the maintainers, currently only PascalSchroederDE. He will then review it cautiously and merge it to master if everything looks fine.

### Future Plans

Right now it is aimed to extend this library to all other features of the LemonMarketAPI, especially implementing the new latest data subscribtion. Also it is focused on keeping it updated according to API changes. Further also all parameters should be able to include, which are possible to be included in API requests as right now only the most important and required ones are usable.

