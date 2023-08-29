package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CacheCleaner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="36.096" cy="22.654" r="1.393" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="40.222" cy="26.814" r="1.672" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.213 10.465l11.304 1.98m-7.97 2.031l7.345 1.302m-9.48 1.562l8.959 1.563m6.579 13.093l-.277-3.404c-.694-4.605-6.012-5.723-7.657-1.276c-1.821 3.764-8.313 7.281-15.418 8.151c2.863 9.246 18.17 7.426 20.263 4.74"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.954 34.79c-4.981 2.828-.245 7.569 2.605 5.73c2.071 2.776 6.469 2.25 7.5.156c5.886 2.626 7.238-5.573 2.397-6.302c-1.076-3.761-4.968-4.413-7.293-2.136c-1.987-.49-4.029-1.266-5.209 2.552ZM9.296 17.407v1.51m0 2.353v1.51m1.229-2.686l1.51.052m-5.373 0h1.51m5.899 7.796l1.031 1.105m0-3.873l-1.068 1.068m-2.731-.995l1.068 1.068m0 1.59l-1.068 1.068m9.713-1.657c2.328 1.94 4.939 1.975 7.657 1.276M10.978 41.135c5.694-1.93 9.219-4.403 11.184-7.282m-.052 4.636c-1.116 1.544-2.471 2.825-4.766 3.878M7.101 38.31c3.242-.09 5.777-1.216 8.238-2.478m8.373-11.216l3.216-17.172c.448-2.972 3.948-2.119 3.269.573L26.833 25.29"/>`),
		g.Group(children),
	)
}