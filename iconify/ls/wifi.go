package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M486 476c-43-42-110-42-153 0c-42 42-42 110 0 153c43 42 110 42 153 0c42-43 42-111 0-153zm-226-73c-28 28-46 61-55 97l-67-65c14-34 35-64 63-91c115-116 302-116 417 0c27 27 49 57 63 91l-67 65c-9-36-27-69-55-97c-83-83-217-83-299 0zM127 270c-26 27-49 56-66 87L0 296c19-30 42-58 68-85c188-188 496-188 684 0c26 27 48 55 67 85l-61 61c-17-31-40-60-66-87c-156-156-409-156-565 0z"/>`),
		g.Group(children),
	)
}