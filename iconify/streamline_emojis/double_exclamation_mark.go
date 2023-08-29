package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M10.41 44.23a13.59 1.77 0 1 0 27.18 0a13.59 1.77 0 1 0-27.18 0Z" opacity=".15"/><path fill="#ff6242" d="M33.15 3.25h-2.74c-1.82 0-3.26 1.15-3.15 2.51l1.62 20.74c.09 1.15 1.36 2.05 2.9 2.05s2.81-.9 2.9-2.05L36.3 5.76c.1-1.36-1.3-2.51-3.15-2.51Z"/><path fill="#ff866e" d="M27.44 8.11a3.26 3.26 0 0 1 3-1.56h2.74a3.24 3.24 0 0 1 3 1.56l.19-2.35C36.4 4.4 35 3.25 33.15 3.25h-2.74c-1.82 0-3.26 1.15-3.15 2.51Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.15 3.25h-2.74c-1.82 0-3.26 1.15-3.15 2.51l1.62 20.74c.09 1.15 1.36 2.05 2.9 2.05s2.81-.9 2.9-2.05L36.3 5.76c.1-1.36-1.3-2.51-3.15-2.51Z"/><path fill="#ff6242" d="M28.22 35.2a3.56 3.56 0 1 0 7.12 0a3.56 3.56 0 1 0-7.12 0Z"/><path fill="#ff866e" d="M31.78 33.93a3.49 3.49 0 0 1 3.48 2a3.67 3.67 0 0 0 .08-.75a3.57 3.57 0 0 0-7.13 0a3.08 3.08 0 0 0 .09.75c.34-1.14 1.77-2 3.48-2Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M28.22 35.2a3.56 3.56 0 1 0 7.12 0a3.56 3.56 0 1 0-7.12 0Z"/><path fill="#ff6242" d="M17.59 3.25h-2.74C13 3.25 11.6 4.4 11.7 5.76l1.62 20.74c.09 1.15 1.37 2.05 2.9 2.05s2.81-.9 2.9-2.05l1.62-20.74c.11-1.36-1.33-2.51-3.15-2.51Z"/><path fill="#ff866e" d="M11.89 8.11a3.24 3.24 0 0 1 3-1.56h2.74a3.26 3.26 0 0 1 3 1.56l.18-2.35c.11-1.36-1.33-2.51-3.15-2.51h-2.81C13 3.25 11.6 4.4 11.7 5.76Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M17.59 3.25h-2.74C13 3.25 11.6 4.4 11.7 5.76l1.62 20.74c.09 1.15 1.37 2.05 2.9 2.05s2.81-.9 2.9-2.05l1.62-20.74c.11-1.36-1.33-2.51-3.15-2.51Z"/><path fill="#ff6242" d="M12.66 35.2a3.56 3.56 0 1 0 7.12 0a3.56 3.56 0 1 0-7.12 0Z"/><path fill="#ff866e" d="M16.22 33.93c1.71 0 3.14.86 3.48 2a3.08 3.08 0 0 0 .09-.75a3.57 3.57 0 0 0-7.13 0a3.67 3.67 0 0 0 .08.75a3.49 3.49 0 0 1 3.48-2Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M12.66 35.2a3.56 3.56 0 1 0 7.12 0a3.56 3.56 0 1 0-7.12 0Z"/>`),
		g.Group(children),
	)
}