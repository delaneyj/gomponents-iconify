package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M336 336h80v80h-80zm-64-64h64v64h-64zm144 144h64v64h-64zm16-144h48v48h-48zM272 432h48v48h-48zm64-336h80v80h-80z"/><path fill="currentColor" d="M480 240H272V32h208Zm-164-44h120V76H316ZM96 96h80v80H96z"/><path fill="currentColor" d="M240 240H32V32h208ZM76 196h120V76H76Zm20 140h80v80H96z"/><path fill="currentColor" d="M240 480H32V272h208ZM76 436h120V316H76Z"/>`),
		g.Group(children),
	)
}