package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterZ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="M108.91 99.71H62.6l48.33-79.77c.44-.73.46-1.64.04-2.39c-.42-.74-1.21-1.2-2.06-1.2H81.06c-.25 0-.49.05-.71.13H19.09c-1.24 0-2.24 1-2.24 2.24v16.72c0 1.24 1.01 2.24 2.24 2.24h47.43l-49.44 79.75c-.45.73-.47 1.65-.06 2.4c.42.75 1.21 1.21 2.07 1.21h29.25c.25 0 .49-.05.72-.12h59.84c1.23 0 2.24-1 2.24-2.24v-16.72c.01-1.25-1-2.25-2.23-2.25z"/>`),
		g.Group(children),
	)
}