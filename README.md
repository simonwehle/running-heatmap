# Activity Heatmap

Display activity routes on a map

## Prerequitsites

### Activities

`activity-heatmap` currently supports only `.gpx` files to for display on the map.
Add your GPX files in the `/activities` folder.

The easiest way to get activities from Garmin is to connect your watch and copy `.fit` files from `/GARMIN/Activity` to your `/activities` folder. For Mac you can use OpenMTP

```sh
brew install --cask openmtp
```

Afterwards you need to convert the `.fit` files to `.gpx`

```bash
brew install gpsbabel
```

```bash
for f in *.fit; do gpsbabel -i garmin_fit -f "$f" -o gpx -F "${f%.fit}.gpx"; done
```

### Map

Open an free [Protomaps Account](https://protomaps.com/account) and create an API token. Set the CORS to `http://localhost:3465`, `https://your.domain.com` or alternatively to `*` to allow all.

## Usage

### Locally

Paste your API token at the end of the style URL

```
MAPLIBRE_STYLE="https://api.protomaps.com/styles/v5/light/en.json?key=MY_KEY" go run .
```

You can view the map via http://localhost:3465

### Docker

Rename the `/docker/.env.example` to `.env` and set your API token.
Start the container with

```
docker compose up -d
```
