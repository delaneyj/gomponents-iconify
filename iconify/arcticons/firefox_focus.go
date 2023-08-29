package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirefoxFocus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.92 43.788A21.602 21.602 0 0 1 24 46.283h0c-5.364 0-21.5-3.667-21.5-22.301a17.75 17.75 0 0 1 8.17-15.014a14.334 14.334 0 0 0 .717 5.458h.071a17.76 17.76 0 0 1 5.124-4.049a12.857 12.857 0 0 0-.262 2.508a11.068 11.068 0 0 0 .322 2.64a4.203 4.203 0 0 1 .49.43a11.093 11.093 0 0 0 2.628 2.27c2.484 1.433 4.24 2.257 4.24 2.257s-1.194 2.389-2.64 2.389c-5.9 0-7.083 3.834-7.083 3.834s1.194 7.931 9.556 7.931c3.081 0 9.173-1.696 9.173-8.827a8.839 8.839 0 0 0-4.623-7.883a7.453 7.453 0 0 1 5.614 1.636a11.42 11.42 0 0 0-9.997-5.9h-.705a15.814 15.814 0 0 1 7.43-11.945a23.893 23.893 0 0 0 3.738 5.71c2.604 2.95 5.972 6.51 7.358 10.177a24.544 24.544 0 0 0-2.066-6.366S45.5 14.606 45.5 24.913a21.618 21.618 0 0 1-1.968 8.722"/><circle cx="38.5" cy="38.5" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.5 34.5h4m-4 4h2.6m-2.6-4v8"/>`),
		g.Group(children),
	)
}