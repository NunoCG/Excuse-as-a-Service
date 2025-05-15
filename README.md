# Excuse-as-a-Service 🐾🚫

“I’m not lazy, I’m just on energy-saving mode.” — Definitely not a developer

## What Is This?

A micro-API that delivers:

- Random developer excuses when your manager asks “Why isn’t that feature live yet?”
- A hard “no” (powered by No-as-a-Service) when your designer says “Can we ship this by EOD?”

All courtesy of caffeine, questionable deadlines, and a dash of cosmic chaos.

## 🚀 Quickstart

1. **Build the image**:

```bash
docker build -t excuse-service:dev .
```

2. **Run the container**:

```bash
docker run -d --name excuse-as-a-service -p 8080:8080 excuse-service:dev
```

3. **Get excuses**:

- **GET** `/excuse` → {"excuse":"My cat deployed to production by accident."}
- **POST** `/excuse` with `{"request":"Yes, please?"}` → {"response":"Nope. Dream on."}

## 🧠 How It Works

1. `excuses.go` picks a random excuse from our lovingly curated list.
2. `main.go` checks your request for “yes.”
    - If found, it fetches a firm “no” from the No-as-a-Service HTTP API.
    - Otherwise, it hands you a whimsical excuse so good you’ll almost feel guilty.

## 🔧 Configuration

- **Port**: Defaults to `8080`; override with `-PORT` env var.
- **Excuses**: Tweak `excuses.go` to add your own classic “the dog ate my code” apologia.

## 🤝 Contributing

- Fork it.
- Add more outrageous excuses.
- Submit a pull request and bask in the glory of knowing you saved humanity from productivity.

## 🎉 Credits & Thanks

- [No-as-a-Service](https://github.com/hotheadhacker/no-as-a-service) for the hard “no.”
- All the developers who’ve ever blamed their cat.

> **__Disclaimer__**: No cats were harmed in the making of these excuses. Actually, we’re not sure. Maybe your mileage may vary.

Enjoy turning every “Can you…?” into “Sorry, can’t.” 😎

<!-- GitAds-Verify: KD3B5RF3PXJJWCU1RXU1FAWQL8C2PM8B -->