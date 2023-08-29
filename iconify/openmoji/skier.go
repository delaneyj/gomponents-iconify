package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiSkier0" stroke-linecap="round" stroke-linejoin="round" d="m38.167 23l.672 3.29c.22 1.078-.292 2.535-1.139 3.237l-2.78 2.307M40 44l1.607 8.04C41.823 53.116 42.675 54 43.5 54s1.5-.9 1.5-2v-5c0-1.1.402-2.805.894-3.79L47 41M11 61l2 2h48m5-5H41m-12 0h-9l-2-2m15-25l23 18m-1-3l-2 3"/></defs><g fill="#FCEA2B"><circle cx="33.969" cy="13.094" r="3"/><path d="M45 25.688L51 35l-2 5l-12 5l.174 14H33.75l-3.281-16.387l5.703-4.913c.952-.766.908-.888.908-1.419c0-.72-.517-2.531-.517-2.531l-3.188-8.125L34.5 20.5l2.938-.938l2.874-.062L45 25.688z"/><path d="m40 44l2 10h3v-9l2-4"/></g><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><circle cx="33.969" cy="13.094" r="3"/><path stroke-linecap="round" stroke-linejoin="round" d="m44.54 24.674l5.4 8.63c.582.933.726 2.532.317 3.553l-.514 1.286c-.409 1.021-1.574 2.203-2.589 2.627l-8.308 3.46c-1.015.424-1.835 1.67-1.822 2.77l.125 10c.014 1.1-.746 2-1.687 2c-.942 0-1.889-.883-2.105-1.96l-2.496-12.466c-.216-1.078.36-2.456 1.279-3.06l4.36-2.87"/><path stroke-linecap="round" stroke-linejoin="round" d="m44.861 25.169l-2.536-3.914c-1.28-1.79-3.9-2.58-5.825-1.755c-1.925.825-2.973 3.221-2.33 5.325l.861 2.227"/><use href="#openmojiSkier0" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.969" cy="13.094" r="3"/><path stroke-linecap="round" stroke-linejoin="round" d="m46.317 27.417l3.632 5.88c.578.937.717 2.539.308 3.56l-.514 1.286c-.409 1.021-1.574 2.203-2.589 2.627l-8.308 3.46c-1.015.424-1.835 1.67-1.822 2.77l.125 10c.014 1.1-.746 2-1.687 2c-.942 0-1.889-.883-2.105-1.96l-2.496-12.466c-.216-1.078.36-2.456 1.279-3.06l4.36-2.87"/><path stroke-linecap="round" stroke-linejoin="round" d="m48.105 30.175l-5.78-8.92c-1.28-1.79-3.9-2.58-5.825-1.755c-1.925.825-2.973 3.221-2.33 5.325l.861 2.227"/><use href="#openmojiSkier0" stroke-linecap="round" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}