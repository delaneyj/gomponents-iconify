package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandWash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M58 461q4 22 21.5 36.5T119 512h274q22 0 39.5-14.5T454 461l58-265q2-8-3-16t-14-9q-9-2-17 3t-9 14l-57 264q-3 17-22 17H119q-18 0-21-17L43 188q-1-8-9.5-13.5T17 171q-8 1-13.5 9T0 196zm177-77q0 21 21 21t21-21V171h-42v213zm64-43q0 22 21 22t21-22V171h-42v170zm-171-85q9 0 15-6l28-28v98q0 21 21 21t21-21V117q0-36 27.5-58T303 45q27 4 43.5 25.5T363 119v180q0 21 21 21t21-21V117q0-51-38-86T277 0q-46 4-76 39.5T171 122v38l-58 58q-6 6-6 14.5t6 14.5q5 9 15 9z"/>`),
		g.Group(children),
	)
}