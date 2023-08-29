package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Racoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 10.848A21.327 21.327 0 0 0 5.202 22.102a5.935 5.935 0 0 0 0 5.601a21.326 21.326 0 0 0 37.596 0a5.935 5.935 0 0 0 0-5.601A21.327 21.327 0 0 0 24 10.848Z"/><ellipse cx="24" cy="32.664" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.212" ry="4.702"/><ellipse cx="24" cy="32.258" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.686" ry="2.033"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.026 28.536c1.35-2.398 1.518-5.388 0-6.815c-1.7-1.595-7.398-1.31-11.046 6.241a14.323 14.323 0 0 0 11.046 8.83"/><circle cx="19.156" cy="24.144" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.674 26.652a17.755 17.755 0 0 1-1.378-2.28s3.306-4.104 7.637-4.104a4.245 4.245 0 0 1 2.633.763M7.233 19.005C6.025 16.721 5.449 12.19 5.41 9.942a1.011 1.011 0 0 1 1.03-1.036c1.937.039 6.227.539 8.377 4.018m12.157 15.612c-1.35-2.398-1.518-5.388 0-6.815c1.7-1.595 7.398-1.31 11.046 6.241a14.323 14.323 0 0 1-11.046 8.83"/><circle cx="28.844" cy="24.144" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.326 26.652a17.755 17.755 0 0 0 1.378-2.28s-3.306-4.104-7.637-4.104a4.245 4.245 0 0 0-2.633.763m12.333-2.026c1.208-2.284 1.784-6.816 1.822-9.063a1.011 1.011 0 0 0-1.028-1.036c-1.937.039-6.227.539-8.377 4.018"/>`),
		g.Group(children),
	)
}