package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M13 0a5 5 0 0 0-5 5v8a5 5 0 0 0 5 5c.173 0 .332-.014.5-.031a5.049 5.049 0 0 0 3.656-2.188c.54-.798.844-1.745.844-2.781V5a5 5 0 0 0-5-5zm0 2c1.3 0 2.396.842 2.813 2H15c-.551 0-1 .448-1 1c0 .965 1 1 1 1h1v1h-1c-.551 0-1 .448-1 1c0 .965 1 1 1 1h1v1h-6V9h1s1-.035 1-1c0-.552-.449-1-1-1h-1V6h1s1-.035 1-1c0-.552-.449-1-1-1h-.813c.417-1.158 1.514-2 2.813-2zm-8 9c-.891 0-1 1-1 1v1c0 4.275 2.998 7.837 7 8.75V24H8c-.959 0-1 1-1 1v1h12v-1c0-.551-.448-1-1-1h-3v-2.25c4.002-.913 7-4.475 7-8.75v-1s-.094-1-1-1s-1 1-1 1v1c0 3.859-3.141 7-7 7s-7-3.141-7-7v-1s-.109-1-1-1z"/>`),
		g.Group(children),
	)
}