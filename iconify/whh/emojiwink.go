package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emojiwink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm-96-640H224q-13 0-22.5 9.5T192 416t9.5 22.5T224 448h192q13 0 22.5-9.5T448 416t-9.5-22.5T416 384zm320-128q-40 0-68 37.5T640 384t28 90.5t68 37.5t68-37.5t28-90.5t-28-90.5t-68-37.5zm64 384q-13 0-22.5 9.5T768 672q-14 20-28.5 35.5t-27 25.5t-31.5 17t-32.5 10.5t-40 5t-43.5 2t-53 .5t-53-.5t-43.5-2t-40-5T343 750t-31.5-17t-27-25.5T256 672q0-13-9.5-22.5T224 640t-22.5 9.5T192 672q24 73 110 116.5T512 832t210-43.5T832 672q0-13-9.5-22.5T800 640z"/>`),
		g.Group(children),
	)
}