package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Svn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 200.093c-146.988 8.553-366.786 48.059-380.439 18.253C127.278 181.962 363.655 117.21 512 97.842v102.251zM0 299.653v110.608c161.463-19.058 423.402-90.818 409.523-126.903C394.83 248.491 144.984 294.122 0 299.654zm493.867-230.41H0v134.72c122.23-49.787 329.9-103.877 493.867-134.72zM26.077 442.757H512V309.12c-139.957 55.892-358.784 113-485.924 133.637z"/>`),
		g.Group(children),
	)
}