package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiftyOneHundredTwenty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M503 21q-9-2-17 3t-9 14l-12 67q-5-3-15-13q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28q-16 15-28 15t-28-15Q84 66 55 64l-4-26q-2-8-10-13t-16-4t-13.5 9.5T8 47l58 286q4 22 21.5 36.5T127 384h274q22 0 39.5-14.5T462 333l58-286q2-8-3.5-16.5T503 21zm-81 303q-3 17-21 17H127q-18 0-21-17L63 109q3 2 7.5 6.5T78 122q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27q35 0 57-27q15-15 28-15t28 15q19 19 36 25zm-137-89q0 8-6 14.5t-15 6.5t-15-6.5t-6-14.5q0-9 6-15.5t15-6.5t15 6.5t6 15.5zm107 0q0 8-6.5 14.5T371 256q-9 0-15.5-6.5T349 235q0-9 6.5-15.5T371 213q8 0 14.5 6.5T392 235zm-213 0q0 8-6.5 14.5T157 256q-8 0-14.5-6.5T136 235q0-9 6.5-15.5T157 213q9 0 15.5 6.5T179 235z"/>`),
		g.Group(children),
	)
}