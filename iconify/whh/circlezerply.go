package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlezerply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512q0-104 40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199q0 139-68.5 257T769 955.5T512 1024zm171-324q-29 0-66-23t-83.5-55t-85.5-46q277-264 277-340q0-11-.5-16.5t-4.5-13t-13-11t-24-3.5q-13 0-34 10t-43 22t-64.5 22t-93.5 10q-19 0-35-16t-22-32l-7-16q-2 1-4.5 3t-10 9.5t-13 16.5t-10.5 25t-5 34q0 32 20 52t55 20t85-9.5t75-22.5q-31 34-93.5 90.5t-108 97.5t-82 89t-36.5 81q0 15 17 18.5t68 3.5q35 0 67 13.5t58 33t53 39t63.5 33T661 832q22 0 39.5-10t28.5-24.5t19-35t11.5-37t6-35T768 666v-10q-3 5-9 12.5T731 688t-48 12z"/>`),
		g.Group(children),
	)
}