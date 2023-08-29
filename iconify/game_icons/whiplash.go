package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whiplash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M407.056 376.454C511.348 517.65 270.95 424.988 194.373 328.11c-59.935-75.823 212.347-41.197 101.407-177.47C172.653-.453-54.165-13.92 65.816 92.687C35.032 6.53 299.823 128.57 311.883 205.35c10.934 69.623-308.9 30.456-112.237 175.655c137.22 101.312 397.83 144.363 207.41-4.55z"/>`),
		g.Group(children),
	)
}