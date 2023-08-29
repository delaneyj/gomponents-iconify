package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarezerply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-213-324q-29 0-66-23t-83.5-55t-85.5-46q277-264 277-340q0-11-.5-16.5t-4-13t-13-11t-24.5-3.5q-13 0-34 10t-43 22t-64.5 22t-93.5 10q-19 0-35-16t-22-32l-7-16q-2 1-5 3t-10 9.5t-12.5 16.5t-10.5 25t-5 34q0 32 20 52t55 20t85-9.5t75-22.5q-31 34-93.5 90.5t-108 97.5t-82 89t-36.5 81q0 15 17 18.5t68 3.5q35 0 67 13.5t58 33t53 39t63.5 33t78.5 13.5q22 0 39.5-10t28.5-24.5t19-35t11.5-37t6-35t2.5-24.5v-10q-3 5-9 12.5t-28 19.5t-48 12z"/>`),
		g.Group(children),
	)
}