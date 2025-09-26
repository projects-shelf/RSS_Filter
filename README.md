# shelf | RSS_Filter

### A text-based RSS feed filter

## Quickstart

```shell
git clone https://github.com/projects-shelf/RSS_Filter.git
cd RSS_Filter
docker-compose up -d
```

Once running, you can access a filtered RSS feed by visiting:

```
http://localhost:8080/?url=https://example.com/rss.xml
```

## Features

### Filters out RSS feed items containing links to domains listed in blocklist.txt.

Example `blocklist.txt`:

```txt
example.com
example.net
example.org
```

## License

RSS_Filter is licensed under [MIT License](https://github.com/projects-shelf/RSS_Generator/blob/main/LICENSE).

## Author

Developed by [PepperCat](https://github.com/PepperCat-YamanekoVillage).
