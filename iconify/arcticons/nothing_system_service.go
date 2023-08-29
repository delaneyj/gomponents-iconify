package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NothingSystemService(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="16" cy="18" r=".75" fill="currentColor"/><circle cx="14" cy="21" r=".75" fill="currentColor"/><circle cx="18" cy="18" r=".75" fill="currentColor"/><circle cx="20" cy="18" r=".75" fill="currentColor"/><circle cx="16" cy="30" r=".75" fill="currentColor"/><circle cx="18" cy="30" r=".75" fill="currentColor"/><circle cx="20" cy="30" r=".75" fill="currentColor"/><circle cx="26" cy="30" r=".75" fill="currentColor"/><circle cx="28" cy="30" r=".75" fill="currentColor"/><circle cx="30" cy="30" r=".75" fill="currentColor"/><circle cx="32" cy="30" r=".75" fill="currentColor"/><circle cx="22" cy="21" r=".75" fill="currentColor"/><circle cx="26" cy="21" r=".75" fill="currentColor"/><circle cx="14" cy="24" r=".75" fill="currentColor"/><circle cx="22" cy="24" r=".75" fill="currentColor"/><circle cx="14" cy="27" r=".75" fill="currentColor"/><circle cx="22" cy="27" r=".75" fill="currentColor"/><circle cx="34" cy="27" r=".75" fill="currentColor"/><circle cx="28" cy="24" r=".75" fill="currentColor"/><circle cx="30" cy="24" r=".75" fill="currentColor"/><circle cx="32" cy="24" r=".75" fill="currentColor"/><circle cx="28" cy="18" r=".75" fill="currentColor"/><circle cx="30" cy="18" r=".75" fill="currentColor"/><circle cx="32" cy="18" r=".75" fill="currentColor"/><circle cx="34" cy="18" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.23 26c.087-.663.134-1.331.14-2a16.516 16.516 0 0 0-.14-2l4.33-3.39a1 1 0 0 0 .25-1.31l-4.1-7.11a1 1 0 0 0-1.25-.44l-5.11 2.06a15.683 15.683 0 0 0-3.46-2l-.77-5.43a1 1 0 0 0-1-.86H19.9a1 1 0 0 0-1 .86l-.77 5.43a15.358 15.358 0 0 0-3.46 2L9.54 9.75a1 1 0 0 0-1.25.44l-4.1 7.11a1 1 0 0 0 .25 1.31L8.76 22a16.66 16.66 0 0 0-.14 2c.006.669.053 1.337.14 2l-4.32 3.39a1 1 0 0 0-.25 1.31l4.1 7.11a1 1 0 0 0 1.25.44l5.11-2.06a15.683 15.683 0 0 0 3.46 2l.77 5.43a1 1 0 0 0 1 .86h8.2a1 1 0 0 0 1-.86l.77-5.43a15.358 15.358 0 0 0 3.46-2l5.11 2.06a1 1 0 0 0 1.25-.44l4.1-7.11a1 1 0 0 0-.25-1.31L39.23 26Z"/>`),
		g.Group(children),
	)
}