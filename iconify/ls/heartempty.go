package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heartempty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M384.022 668c5 0 11-1 15-4c11-7 259-165 344-320c36-66 32-155-9-224c-36-61-96-95-168-95c-99 0-153 53-182 98c-29-45-84-98-182-98c-73 0-132 34-169 95c-41 69-44 158-8 224c84 155 333 313 343 320c5 3 10 4 16 4zm-182-583c118 0 148 98 152 117c4 13 16 24 30 24c13 0 26-11 29-24c5-19 35-117 153-117c65 0 100 35 117 66c30 49 33 118 8 164c-67 120-252 250-307 288c-56-38-241-168-307-288c-25-46-22-115 8-164c17-31 52-66 117-66z"/>`),
		g.Group(children),
	)
}