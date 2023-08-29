package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenMagO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M50.71 22.467a24.908 24.908 0 0 0-1.136.01c-19.345-.536-33.704 23.186-24.09 40.07c7.63 15.658 31.001 20.382 43.461 7.734c12.163-11.063 11.582-32.995-2.386-42.357c-4.557-3.38-10.166-5.377-15.848-5.457zm-.565 5.105a22.36 22.36 0 0 1 15.863 6.59a22.386 22.386 0 0 1 0 31.727a22.384 22.384 0 0 1-31.725 0a22.386 22.386 0 0 1 0-31.727a22.359 22.359 0 0 1 15.862-6.59z" color="currentColor"/><path fill="currentColor" d="M5 17.5h90v65H5Zm-2.5-5A2.5 2.5 0 0 0 0 15v70a2.5 2.5 0 0 0 2.5 2.5h95A2.5 2.5 0 0 0 100 85V15a2.5 2.5 0 0 0-2.5-2.5z" color="currentColor"/>`),
		g.Group(children),
	)
}