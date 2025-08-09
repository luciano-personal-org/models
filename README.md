# Algorithmic Trading Models

A comprehensive Go package defining data structures for algorithmic trading operations, market data processing, and order management.

## Overview

This package provides type-safe models for:
- **Trading Operations**: Trades, orders, and bookings
- **Market Data**: Candles, quotes, and order books
- **Order Management**: Order types, execution tracking, and broker integration
- **Message Processing**: Trading platform message types and connectors

## Installation

```bash
go get github.com/luciano-personal-org/models
```

## Models

### Core Trading Models

#### Trade
Represents a complete trading operation with entry and exit orders:

```go
type Trade struct {
    ID          string      `json:"ID"`
    FirstOrder  FirstOrder  `json:"FirstOrder"`  // Entry order
    SecondOrder SecondOrder `json:"SecondOrder"` // Exit order
    Quantity    float64     `json:"Quantity"`
    Candle      Candle      `json:"Candle"`      // Market data context
    Gain        float64     `json:"Gain"`        // Trading result
}
```

#### Orders
Container for order pairs with market context:

```go
type Orders struct {
    ID       string      `json:"ID"`
    First    FirstOrder  `json:"FirstOrder"`
    Second   SecondOrder `json:"SecondOrder"`
    Strength float64     `json:"Strength"`    // Signal strength (0.0-1.0)
    Candle   Candle      `json:"Candle"`
    Volume   float64     `json:"Volume"`
}
```

### Market Data Models

#### Candle
OHLCV candlestick data with metadata:

```go
type Candle struct {
    ID              string  `json:"ID"`
    Symbol          string  `json:"Symbol"`
    TimeStamp       string  `json:"TimeStamp"`
    Date            string  `json:"Date"`
    Current         float64 `json:"Current"`
    Open            float64 `json:"Open"`
    High            float64 `json:"High"`
    Low             float64 `json:"Low"`
    Close           float64 `json:"Close"`
    Volume          uint64  `json:"Volume"`
    VolumeVariation float64 `json:"VolumeVariation"`
}
```

#### Book
Order book entries with position and pricing:

```go
type Book struct {
    Id        string  `json:"ID"`
    StockName string  `json:"StockName"`
    BookType  byte    `json:"BookType"`
    Position  uint    `json:"Position"`
    Direction byte    `json:"Direction"`
    Price     float64 `json:"Price"`
    DateTime  string  `json:"DateTime"`
    OrderId   string  `json:"OrderId"`
    Quantity  uint    `json:"Quantity"`
    OfferType byte    `json:"OfferType"`
    Average   uint64  `json:"Average"`
}
```

### Order Management

#### Order Types
```go
type OrderType int

const (
    LIMIT OrderType = iota
    MARKET
    STOP_LOSS
    STOP_LOSS_LIMIT
    TAKE_PROFIT
    TAKE_PROFIT_LIMIT
    LIMIT_MAKER
)
```

#### Order Sides
```go
type OrderSide int

const (
    BUY OrderSide = iota
    SELL
)
```

### Booking System

#### Booking (v1)
Original booking structure:

```go
type Booking struct {
    ID                     string          `json:"ID"`
    Orders                 Orders          `json:"Orders"`
    Quantity               float64         `json:"Quantity"`
    QuantityAfterComission float64         `json:"QuantityAfterComission"`
    MarketDirection        MarketDirection `json:"MarketDirection"`
}
```

#### Bookingv2
Updated booking with trade arrays:

```go
type Bookingv2 struct {
    ID              string          `json:"ID"`
    Trades          []Trade         `json:"Trades"`
    MarketDirection MarketDirection `json:"MarketDirection"`
}
```

### Execution Tracking

#### ExecutedOrders
Tracks order execution from broker:

```go
type ExecutedOrders struct {
    ID           string       `json:"ID"`
    Orders       Orders       `json:"Orders"`
    BrokerOrders BrokerOrders `json:"BrokerOrders"`
}
```

### Message Types

#### Market Data Messages
```go
type MessageType byte

const (
    QuoteType MessageType = 'T'  // Quote Messages
    BookType  MessageType = 'B'  // Book Messages
    SynType   MessageType = 'S'  // News Messages
)
```

#### Book Operations
```go
const (
    BookAdd    byte = 'A'  // Add book entry
    BookUpdate byte = 'U'  // Update book entry
    BookEnd    byte = 'E'  // End of messages
    BookCancel byte = 'D'  // Cancel book entry
    BookBuy    byte = 'A'  // Buy direction
    BookSell   byte = 'V'  // Sell direction
)
```

## Usage Examples

### Creating a Trade

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/luciano-personal-org/models"
)

func main() {
    trade := models.Trade{
        ID:       "trade-001",
        Quantity: 100.0,
        Gain:     15.50,
        FirstOrder: models.FirstOrder{
            ID:        "order-001",
            OrderSide: models.BUY,
            OrderType: models.MARKET,
            Price:     50.25,
        },
        SecondOrder: models.SecondOrder{
            ID:        "order-002",
            StopGain:  55.00,
            OrderSide: models.SELL,
            OrderType: models.LIMIT,
        },
        Candle: models.Candle{
            Symbol: "AAPL",
            Open:   50.00,
            High:   52.00,
            Low:    49.50,
            Close:  51.00,
            Volume: 1000000,
        },
    }

    // Serialize to JSON
    data, _ := json.MarshalIndent(trade, "", "  ")
    fmt.Println(string(data))
}
```

### Working with Market Direction

```go
direction := models.BULLISH
fmt.Printf("Market Direction: %s\n", direction.String()) // Output: BULLISH

// Check direction
if direction == models.BULLISH {
    fmt.Println("Going long!")
}
```

### Processing Order Types

```go
orderType := models.STOP_LOSS
fmt.Printf("Order Type: %s\n", orderType.String()) // Output: STOP_LOSS

// Create different order types
orders := []models.OrderType{
    models.LIMIT,
    models.MARKET,
    models.STOP_LOSS,
}

for _, order := range orders {
    fmt.Printf("Processing %s order\n", order.String())
}
```

## Development

### Building
```bash
go build
```

### Testing
```bash
go test ./...
```

### Code Quality
```bash
go fmt ./...
go vet ./...
```

### Module Management
```bash
go mod tidy     # Clean dependencies
go mod verify   # Verify checksums
```

## Features

- **Type Safety**: All enums implement `String()` methods for safe serialization
- **JSON Ready**: All structs have proper JSON tags for API integration
- **Extensible**: Clear separation between v1 and v2 models for evolution
- **OpenTelemetry**: Built-in tracing support with `CandleV2`
- **Broker Integration**: Models for tracking execution across different brokers

## Architecture

This package follows a pure data model approach:
- No business logic, only data structures
- Immutable constants for message types and operations
- Type-safe enums with string representations
- Clear separation between different domains (trading, market data, execution)

## Dependencies

- `go.opentelemetry.io/otel` - For distributed tracing support
- Standard library only

## License

See LICENSE file for details.
