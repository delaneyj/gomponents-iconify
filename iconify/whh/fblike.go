package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fblike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928 960H576l-64-64H320v64H0V384h320v64h64l96-96L640 0h192v160l-64 96v64h160l96 64v512zM256 448H64v448h192V448zm704-21l-64-43H704V224l64-96V64h-79L544 384L416 512h-96v320h224l64 64h288l64-43V427z"/>`),
		g.Group(children),
	)
}