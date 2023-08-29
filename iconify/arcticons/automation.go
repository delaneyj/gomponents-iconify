package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Automation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 19.77a11.13 11.13 0 0 0-2.26.94l.26 4.53L5.44 25a11.13 11.13 0 0 0-.94 2.26l3.37 3.05l-3.37 3a10.68 10.68 0 0 0 .94 2.26l4.56-.18l-.23 4.54a11.09 11.09 0 0 0 2.26.93l3.05-3.37l3 3.37a10.64 10.64 0 0 0 2.26-.93l-.23-4.54l4.54.23a10.64 10.64 0 0 0 .93-2.26l-3.37-3l3.37-3.05a11.09 11.09 0 0 0-.92-2.31l-4.54.23l.23-4.53a10.68 10.68 0 0 0-2.26-.94l-3 3.37Z"/><circle cx="15.05" cy="30.32" r="3.96" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.91 7.14a11 11 0 0 0-2.27.93l.24 4.54l-4.54-.23a10.64 10.64 0 0 0-.93 2.26l3.36 3l-3.36 3.05a11.09 11.09 0 0 0 .93 2.31l4.54-.23l-.24 4.53a11 11 0 0 0 2.27.94l3-3.37l3 3.37a11.13 11.13 0 0 0 2.26-.94L38 22.76l4.53.23a11.13 11.13 0 0 0 .94-2.26l-3.37-3.05l3.37-3a10.68 10.68 0 0 0-.94-2.26l-4.53.19l.23-4.54A11 11 0 0 0 36 7.14l-3 3.37Z"/><circle cx="32.95" cy="17.68" r="3.96" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}