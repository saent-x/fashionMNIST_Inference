# Fashion-MNIST Inference

A small Go web app that classifies clothing images with a neural network I wrote from scratch in Go. Upload a PNG and the server converts it to a 28×28 grayscale tensor, runs it through the trained Fashion-MNIST model and returns the predicted class (T-shirt, trouser, sneaker and so on).

The network itself (layers, activations, training and model serialisation on top of [gonum](https://www.gonum.org/)) lives in [neural-network-module](https://github.com/saent-x/neural-network-module), which I also used for CAN bus intrusion detection. This repo is the inference side: a saved model plus a server-rendered UI.

## How it works

```
PNG upload ──► /upload handler ──► grayscale + resize to 28×28 ──► scale to [0,1]
                                                                      │
                               trained model (core/saved_models/*.json)
                                                                      ▼
                                               forward pass ──► predicted class
```

## Stack

Go · gonum · [gomponents](https://www.gomponents.com/) for server-side HTML · Tailwind CSS

## Run it

```bash
npm install && make build   # builds the Tailwind stylesheet
make start                  # serves on http://localhost:3000
```

## Layout

| Path | What |
| --- | --- |
| `main.go` | HTTP server and routes |
| `handlers/upload.go` | Image upload and prediction endpoint |
| `core/core.go` | Preprocessing and model inference |
| `core/saved_models/` | Trained Fashion-MNIST model |
| `pages/`, `components/` | gomponents UI |
