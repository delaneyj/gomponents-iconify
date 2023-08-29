package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radioactive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1023q-104 0-199-40.5t-163.5-109T40.5 710T0 511t40.5-198.5t109-163T313 40.5T512 0t199 40.5t163.5 109t109 163T1024 511t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-128V639q-34 0-63-17t-46-45L183 709q51 85 138.5 135.5T512 895zm0-768q-103 0-190.5 51T183 314l220 132q17-29 46-46t63-17V127zm-.5 320q-26.5 0-45 19T448 511.5t18.5 45t45 18.5t45.5-18.5t19-45t-19-45.5t-45.5-19zM841 314L621 446q19 31 19 65t-18 64l222 128q52-89 52-192q0-106-55-197z"/>`),
		g.Group(children),
	)
}