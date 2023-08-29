package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Painting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0q112 0 152 27.5T704 128q0 32-15 61t-33 47t-33 40.5t-15 43.5q0 96 96 96q21 0 43.5-15t40.5-33t47-33t61-15q72 0 100 41t28 151q0 104-40.5 199t-109 163.5t-163.5 109t-199 40.5zM288 832q40 0 68-28t28-68t-28-68t-68-28t-68 28t-28 68t28 68t68 28zm-64-512q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm192-192q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm192 576q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm192-192q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28z"/>`),
		g.Group(children),
	)
}