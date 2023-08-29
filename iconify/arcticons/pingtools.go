package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pingtools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><rect width="31.878" height="8.539" x="8.061" y="9.079" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.264"/><circle cx="35.746" cy="13.348" r=".75" fill="currentColor"/><circle cx="32.092" cy="13.348" r=".75" fill="currentColor"/><circle cx="28.438" cy="13.348" r=".75" fill="currentColor"/><circle cx="13.76" cy="13.348" r="2.911" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="31.878" height="8.539" x="8.061" y="19.731" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.264"/><circle cx="35.746" cy="23.999" r=".75" fill="currentColor"/><circle cx="32.092" cy="23.999" r=".75" fill="currentColor"/><circle cx="28.438" cy="23.999" r=".75" fill="currentColor"/><circle cx="13.76" cy="23.999" r="2.911" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="31.878" height="8.539" x="8.061" y="30.382" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.264"/><circle cx="35.746" cy="34.651" r=".75" fill="currentColor"/><circle cx="32.092" cy="34.651" r=".75" fill="currentColor"/><circle cx="28.438" cy="34.651" r=".75" fill="currentColor"/><circle cx="13.76" cy="34.651" r="2.911" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}