package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BbOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#00267f" d="M0-.2h512V512H0z"/><path fill="#ffc726" d="M170.7-.2h170.6V512H170.7z"/><path id="flagBb1x10" fill="#000" d="M256 173.3c-5.5 15.1-11.2 30.9-23.3 43a51.7 51.7 0 0 1 14.6-2.3v63.6l-18 2.7c-.7 0-.9-1-.9-2.4a243.6 243.6 0 0 0-11.7-53.6c-.4-2.3-7.2-11.3-2-9.7c.7 0 7.7 3 6.6 1.6a68 68 0 0 0-37.1-19.2c-1.2-.3-2 .3-.9 1.7c18 27.7 33.1 60.4 33 99.2c7 0 24-4.1 31-4.1v44.9h8.8l2-125.4z"/><use width="100%" height="100%" href="#flagBb1x10" transform="matrix(-1 0 0 1 512 0)"/>`),
		g.Group(children),
	)
}