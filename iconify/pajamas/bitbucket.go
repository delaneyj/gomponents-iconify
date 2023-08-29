package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitbucket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1.454 1.816a.452.452 0 0 0-.345.153a.435.435 0 0 0-.103.358L2.91 13.684c.049.287.3.498.596.5h9.135a.447.447 0 0 0 .449-.37L14.994 2.33a.435.435 0 0 0-.1-.358a.452.452 0 0 0-.342-.156H1.454Zm8.018 8.208H6.557l-.79-4.05h4.412l-.707 4.05Z"/>`),
		g.Group(children),
	)
}