package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TerraformingMars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="32.078" cy="6.337" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.622" ry="2.288"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.456 6.337v4.078c0 1.264 2.07 2.289 4.622 2.289s4.621-1.025 4.621-2.289V6.337"/><ellipse cx="19.558" cy="14.39" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.622" ry="2.288"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.936 14.39v4.68c0 1.265 2.069 2.29 4.622 2.29s4.621-1.025 4.621-2.29v-4.68"/><ellipse cx="16.159" cy="5.192" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.622" ry="2.288"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.537 5.192v2.875c0 1.264 2.07 2.288 4.622 2.288s4.622-1.024 4.622-2.288V5.192M6.697 11.236a7.858 7.858 0 0 0-.555 2.886c0 6.418 7.995 11.622 17.858 11.622s17.858-5.204 17.858-11.622c0-.996-.192-1.963-.554-2.886"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.537 6.479C6.067 10.377 2.5 16.772 2.5 24c0 11.874 9.626 21.5 21.5 21.5S45.5 35.874 45.5 24c0-7.14-3.48-13.468-8.838-17.378m-4.636-2.574A21.445 21.445 0 0 0 24 2.5a21.49 21.49 0 0 0-5.517.715"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.698 16.851h3.085v2.165h-3.085zm0 2.165h3.085v2.569h-3.085zm-3.242-3.708h3.085v2.165h-3.085zm0 2.165h3.085v4.111h-3.085z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.059 15.946c.908 6.539 8.546 11.732 18.126 12.32v1.666h3.989v-1.687c9.411-.695 16.87-5.842 17.767-12.299"/>`),
		g.Group(children),
	)
}