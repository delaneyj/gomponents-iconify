package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speedalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zM340 509l87 228q-43 38-43 95q0 53 37.5 90.5T512 960t90.5-37.5T640 832q0-50-34-86.5T522 705L407 476q-1-2-4-7q-9-19-15-30.5T371.5 410T351 387t-18 0q-12 9-13 33t5.5 47.5T340 509zM512 64q-91 0-174 35.5T195 195T99.5 338T64 512h2q8-106 65-194t149-139t200-51t200 51t149 139t65 194h66q0-91-35.5-174T829 195T686 99.5T512 64z"/>`),
		g.Group(children),
	)
}