package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<mask id="flagpackMt0" fill="#fff"><path fill-rule="evenodd" d="M10 4H6v2.993A1 1 0 0 0 5.232 8H2v4h3.33c.13.267.375.469.67.539V16h4v-3.449a1 1 0 0 0 .733-.551H14V8h-3.169v-.034A1 1 0 0 0 10 6.98V4Z" clip-rule="evenodd"/></mask><g fill="none"><path fill="#fff" d="M0 0h32v24H0z"/><path fill="#A0A0A0" fill-rule="evenodd" d="M10 4H6v2.993A1 1 0 0 0 5.232 8H2v4h3.33c.13.267.375.469.67.539V16h4v-3.449a1 1 0 0 0 .733-.551H14V8h-3.169v-.034A1 1 0 0 0 10 6.98V4Z" clip-rule="evenodd"/><path fill="#FECA00" d="M6 4V3H5v1h1Zm4 0h1V3h-1v1ZM6 6.993l.23.973l.77-.183v-.79H6ZM5.232 8v1h1.035l-.035-1.034l-1 .034ZM2 8V7H1v1h1Zm0 4H1v1h1v-1Zm3.33 0l.901-.435L5.958 11h-.627v1Zm.67.539h1v-.791l-.77-.182l-.23.973ZM6 16H5v1h1v-1Zm4 0v1h1v-1h-1Zm0-3.449l-.167-.986l-.833.142v.844h1Zm.733-.551v-1h-.628l-.273.565l.9.435ZM14 12v1h1v-1h-1Zm0-4h1V7h-1v1Zm-3.169 0l-1-.034L9.798 9h1.034V8ZM10 6.98H9v.844l.833.142L10 6.98ZM6 5h4V3H6v2Zm1 1.993V4H5v2.993h2Zm-.768.973H6.23L5.769 6.02a2 2 0 0 0-1.537 1.946h2Zm0 0h-2v.068l2-.068ZM2 9h3.232V7H2v2Zm1 3V8H1v4h2Zm2.33-1H2v2h3.33v-2Zm.9.566l.001-.001l-1.8.87a2 2 0 0 0 1.338 1.077l.462-1.946ZM7 16v-3.461H5V16h2Zm3-1H6v2h4v-2Zm-1-2.449V16h2v-3.449H9Zm1.167.986c.65-.11 1.19-.531 1.466-1.102l-1.8-.87l.334 1.972ZM14 11h-3.267v2H14v-2Zm-1-3v4h2V8h-2Zm-2.169 1H14V7h-3.169v2Zm-1-1.034l2 .068v-.068h-2Zm.002 0c-.003 0-.002-.002-.001 0c0-.002 0-.002 0 0h2a2 2 0 0 0-1.665-1.972l-.334 1.972ZM9 4v2.98h2V4H9Z" mask="url(#flagpackMt0)"/><path fill="#E31D1C" fill-rule="evenodd" d="M16 0h16v24H16V0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}