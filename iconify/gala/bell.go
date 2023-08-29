package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaBell0"><g id="galaBell1" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="4.233" transform="translate(16) scale(3.77953)"><path id="galaBell2" stroke-opacity="1" d="M 33.866678,59.266666 A 4.2333331,4.2333331 0 0 1 29.633345,63.5"/><path id="galaBell3" stroke-opacity="1" d="M -25.4,59.266666 A 4.2333331,4.2333331 0 0 1 -29.633333,63.5" transform="scale(-1 1)"/><path id="galaBell4" stroke-opacity="1" d="m 25.400031,55.033482 v 4.233333"/><path id="galaBell5" stroke-opacity="1" d="m 33.866698,55.033482 -1.9e-5,4.233332"/><path id="galaBell6" d="m 55.033333,50.8 c -8.466667,4.233333 -21.166667,4.233333 -25.4,4.233333"/><path id="galaBell7" stroke-opacity="1" d="m 55.033333,50.8 c 0,-8.466667 -8.466667,-4.233336 -8.466667,-25.400002 0,-9.466021 -8.466666,-16.9330112 -12.699988,-16.9330112"/><path id="galaBell8" d="M 4.233334,50.8 C 12.7,55.033333 25.4,55.033333 29.633333,55.033333"/><path id="galaBell9" stroke-opacity="1" d="m 4.2333451,50.799681 c 0,-8.466669 8.4666669,-4.233338 8.4666669,-25.400003 0,-9.46602 8.466666,-16.9330108 12.699988,-16.9330108"/><path id="galaBella" stroke-opacity="1" d="m -25.400012,-8.4666872 a 4.2333331,4.2333331 0 0 1 -4.233333,4.2333331" transform="scale(-1)"/><path id="galaBellb" stroke-opacity="1" d="m 33.866666,-8.4666672 a 4.2333331,4.2333331 0 0 1 -4.233333,4.2333331" transform="scale(1 -1)"/></g></g>`),
		g.Group(children),
	)
}