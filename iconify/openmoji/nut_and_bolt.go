package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NutAndBolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#9b9b9a" d="m36.804 28.439l7.127 7.058l3.03-3.03l-7.092-7.093l-3.065 3.065z"/><path fill="#d0cfce" d="m42.744 11.987l16.934 16.934l-8.131 8.131l-16.934-16.934zM32.174 33.07l7.695-7.696l7.092 7.093l-27.546 27.546l-7.093-7.092L32.174 33.07"/><path fill="#9b9b9a" d="M33.006 32.236L40.1 39.33l-1.36 1.36l-7.093-7.092zm-9.522 9.523l7.092 7.092l-1.36 1.36l-7.092-7.092zm-2.04 2.041l7.092 7.092l-1.36 1.36l-7.092-7.092zm-2.041 2.039l7.093 7.092l-1.36 1.36l-7.093-7.092z"/><path fill="#d0cfce" d="m26.722 31.428l14.185 14.186l-5.32 5.32l-14.184-14.186z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m42.744 11.987l16.934 16.934l-8.131 8.131l-16.934-16.934z"/><path d="m32.755 32.488l7.114-7.114l7.092 7.093l-7.078 7.079m-10.241 10.24l-9.086 9.087l-2.029.253l-5.352-5.352l-.001-1.705l9.201-9.201m4.347-11.44l14.185 14.186l-5.32 5.32l-14.184-14.186z"/></g>`),
		g.Group(children),
	)
}