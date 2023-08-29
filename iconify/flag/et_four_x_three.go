package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EtFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagEt4x30"><path fill-opacity=".7" d="M-61.3 0h682.7v512H-61.3z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagEt4x30)" transform="translate(57.5) scale(.94)"><path fill="#ffc621" d="M-238 3.5H800v498H-238z"/><path fill="#ef2118" d="M-240 342.5H799.3V512H-240z"/><path fill="#298c08" d="M-238 0H800v180H-238z"/><circle cx="534.2" cy="353" r="199.7" fill="#006bc6" transform="matrix(.54 0 0 .54 -25.8 74)"/><path fill="#ffc621" d="m214.3 188.2l-6.5 4.5l23.5 33l6.3-4l-23.4-33.5zm29.4 78l-9.7-6.8l4-12.7l-48.1.7l-14-10.7l65.7-.7l12.2-36.9l6.6 15l-16.7 52zm76.5-70.7l-6.3-4.8l-24.3 32.4l5.6 4.7l25-32.3zM254.8 247l3.5-11.2h13.3L256.4 190l6-16.5l20.5 62.4l38.8.5l-12.2 10.7l-54.7-.2zm90.6 51.2l2.7-7.4l-38.3-13.3l-2.8 7l38.4 13.7zm-69.1-46.4l11.7-.1l4.1 12.6l38.8-28.5l17.6.6l-53.1 38.7l11.5 37.2l-14-8.4l-16.6-52.1zm-19.8 102l7.9.2l.3-40.5l-7.4-.5l-.8 40.9zm22-80.3l3.8 11.1l-10.7 8l39.4 27.7l5 16.8l-53.6-38l-31.5 22.7l3.5-16l44-32.3zm-103.3 13l2.3 7.5l38.7-12.2l-2-7.2l-39 11.9zm83.2-4l-9.4 7.1l-10.8-7.7l-14.2 46l-14.4 10l19.5-62.7l-31.4-23l16.3-1.6l44.4 31.9z"/></g>`),
		g.Group(children),
	)
}