package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HtmlThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.195 2l1.602 17.994L11.99 22l7.212-2.013L20.805 2Zm14.281 4.123l-.534 5.994l.002.033l-.002.074l-.38 4.19l-.041.373L12 18.037l-.004.004l-4.512-1.258l-.306-3.465H9.39l.157 1.762l2.453.665l2.461-.674l.26-2.869H9.577l-.044-.485l-.101-1.136l-.052-.61h5.538l.202-2.232H6.682l-.044-.485l-.1-1.137l-.053-.61h11.044Zm0 0"/>`),
		g.Group(children),
	)
}