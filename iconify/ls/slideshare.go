package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slideshare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M634 66v264c3-2 5-4 8-6c20-14 34 5 21 24c-24 29-70 66-141 95c26 88 12 151-17 190c-17 24-40 38-63 43v1c-50 13-100-11-98-69c0-1-1-71-1-124c-5-1-11-3-18-5v129c4 130-253 90-179-165c-70-29-116-66-141-95c-12-19 1-38 22-24c2 2 5 4 8 6V66C35 30 62 0 96 0h477c34 0 61 30 61 66zm-33 283V96c0-43-14-61-54-61H125c-42 0-54 15-54 61v256c90 47 167 38 209 37c18-1 30 3 36 10c2 1 3 2 4 4c8 7 16 13 23 19c2-21 14-35 45-33c43 1 122 10 213-40zM244 204c45 0 81 35 81 77s-36 77-81 77c-46 0-82-35-82-77s36-77 82-77zm190 0c45 0 82 35 82 77s-37 77-82 77s-82-35-82-77s37-77 82-77z"/>`),
		g.Group(children),
	)
}