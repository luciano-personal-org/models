# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go models package for an algorithmic trading platform. It defines core data structures for trading operations including orders, trades, bookings, candles, and market data.

## Architecture

This package follows a simple models-only architecture:
- Pure Go structs with JSON tags for serialization
- Type-safe enums using iota constants with String() methods
- No external dependencies beyond OpenTelemetry for observability

## Development Commands

### Building and Testing
```bash
go build          # Compile the package
go test           # Run tests (when available)
go mod tidy       # Clean up module dependencies
go fmt ./...      # Format all Go files
go vet ./...      # Static analysis
```

### Module Management
```bash
go mod init       # Initialize module (already done)
go mod download   # Download dependencies
go get <package>  # Add new dependency
```

## Core Models Structure

### Trading Models
- **Trade**: Represents a complete trade with first/second orders, quantity, candle data, and gain
- **Orders**: Container for first and second orders with strength, candle, and volume data
- **FirstOrder/SecondOrder**: Specific order details with side, type, and price information

### Market Data
- **Book**: Market book entries with position, direction, price, and quantity
- **Candle**: OHLCV candlestick data
- **Quote**: Market quote information

### Booking System
- **Booking**: Original booking structure with orders and market direction
- **Bookingv2**: Updated version using Trade arrays instead of Orders
- **MarketDirection**: BULLISH/BEARISH enum for market sentiment

### Order Management
- **OrderSide**: BUY/SELL enum
- **OrderType**: LIMIT, MARKET, STOP_LOSS, etc. enum
- **TypeInForce**: GTC (Good Till Cancelled) enum
- **NewOrderRespType**: ACK, FULL, RESULT enum

All enums implement String() methods for easy serialization and debugging.