package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Candycane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M712.5 1004q-20.5 20-49 20t-48.5-20L107 499l68 67l147-48l-55-55l446 444q20 20 20 48.5t-20.5 48.5zM419 615l-146 49l73 73l147-49zm171 170l-146 49l73 73l147-49zm-36.5-517q-17.5 17-41 18T470 273l-2 2l-24-25l49-145l-25-25l98 97l-7 8q14 18 13 42t-18.5 41zM370 178q-32-33-80-39L222 5q104-19 195 35zm-235 87L0 231q14-88 77-151q21-21 53-40l62 123q-11 9-17 14q-36 36-40 88zm113 180l-146 48l-25-24Q18 410 2 329l173 43z"/>`),
		g.Group(children),
	)
}