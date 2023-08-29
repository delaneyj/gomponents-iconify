package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BiOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagBi1x10"><path fill="gray" d="M60.8 337h175v175h-175z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagBi1x10)" transform="translate(-178 -986) scale(2.9257)"><path fill="#18b637" d="m0 337l146.6 87.5L0 512zm293.1 0l-146.5 87.5L293 512z"/><path fill="#cf0921" d="m0 337l146.6 87.5L293 337zm0 175l146.6-87.5L293 512z"/><path fill="#fff" d="M293.1 337h-27.3L0 495.7V512h27.3l265.8-158.7z"/><path fill="#fff" d="M197.2 424.5a50.6 50.6 0 1 1-101.2 0a50.6 50.6 0 0 1 101.2 0z"/><path fill="#fff" d="M0 337v16.3L265.8 512h27.3v-16.3L27.3 337z"/><path fill="#cf0921" stroke="#18b637" stroke-width="1pt" d="m156.5 405.4l-6.6.1l-3.4 5.6l-3.4-5.6l-6.5-.1l3.2-5.8l-3.2-5.7l6.6-.2l3.4-5.6l3.4 5.7h6.5l-3.1 5.8zm-22 38.2h-6.6l-3.4 5.7l-3.4-5.6l-6.6-.2l3.2-5.7l-3.1-5.8l6.5-.1l3.4-5.6l3.4 5.6l6.6.2l-3.2 5.7zm44.6 0h-6.6l-3.4 5.7l-3.4-5.6l-6.5-.2l3.1-5.7l-3.1-5.8l6.6-.1l3.4-5.6l3.4 5.6l6.5.2l-3.2 5.7z"/></g>`),
		g.Group(children),
	)
}