package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreenStash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.568 14.638a3.594 3.594 0 1 1 4.905.036"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.485 12.009a10.963 10.963 0 0 0-9.883 10.336c-.988 9.793 5.522 18.044 5.522 18.044h4.778a15.812 15.812 0 0 0 1.832-2.33h6.22a9.893 9.893 0 0 0 1.831 2.33h4.779a24.83 24.83 0 0 0 1.795-2.775s8.297-2.004 8.297-7.661H43.5V22h-2.91a12.352 12.352 0 0 0-7.998-9.611A5.873 5.873 0 0 0 26.7 7.612c-.761 0-.343 1.496-.343 3.935c0 .417.014.394-.901.408"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.078 22.708s.19-1.496 2.593-1.496s2.594 1.496 2.594 1.496M15.03 14.729h10.064m-8.161 11.787s-.248 3.165 3.69 3.165s3.69-2.33 3.69-2.33c0-3.01-7.38-1.578-7.38-4.597c0 0-.249-2.421 3.69-2.421s3.69 3.164 3.69 3.164M20.65 31.05V18.81"/>`),
		g.Group(children),
	)
}