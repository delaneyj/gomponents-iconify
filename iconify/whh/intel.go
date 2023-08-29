package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Intel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 836V670q60-49 94-115t34-139q0-125-91-224.5T576 64V0q125 14 227.5 76.5t161.5 161t59 210.5q0 122-69.5 225T768 836zm-288-68q-40 0-68-28t-28-68V288q0-13 9.5-22.5T416 256h64q13 0 22.5 9.5T512 288v448q0 13-9.5 22.5T480 768zm0-640h-64q-13 0-22.5-9.5T384 96V32q0-13 9.5-22.5T416 0h64q13 0 22.5 9.5T512 32v64q0 13-9.5 22.5T480 128zM64 480q0 113 55.5 209T271 840.5T480 896q121 0 224-66v156q-93 38-192 38q-139 0-257-68.5T68.5 769T0 512q0-142 72.5-261.5T265 64h55v64q-119 49-187.5 138T64 480z"/>`),
		g.Group(children),
	)
}