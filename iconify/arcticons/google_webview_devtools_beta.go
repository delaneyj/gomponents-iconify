package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleWebviewDevtoolsBeta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.58 32.815l1.218-2.111c.256-.44.144-.994-.246-1.312l-4.325-3.393c.082-.656.143-1.322.143-1.999s-.061-1.342-.143-1.998l4.335-3.393a1.02 1.02 0 0 0 .246-1.312l-4.1-7.103c-.256-.441-.789-.626-1.25-.441l-5.105 2.06a15.556 15.556 0 0 0-3.464-2.019l-.769-5.433a1.054 1.054 0 0 0-1.025-.86h-8.2c-.512 0-.933.379-1.015.86l-.768 5.433a15.167 15.167 0 0 0-3.465 2.02L9.543 9.752a1.025 1.025 0 0 0-1.25.44l-4.1 7.104c-.257.44-.144.994.245 1.312l4.326 3.393c-.09.662-.138 1.33-.144 1.999c0 .676.062 1.342.144 1.998l-4.326 3.393c-.39.307-.502.861-.246 1.312l4.1 7.103c.256.441.79.625 1.25.441l5.105-2.06a15.56 15.56 0 0 0 3.465 2.019l.768 5.433c.082.481.503.86 1.015.86h8.2c.513 0 .933-.379 1.015-.86l.769-5.433c.59-.246 1.163-.529 1.716-.847"/><circle cx="23.995" cy="24" r="9.264" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.995 14.736h18.14M31.012 30.057L19.159 44.182M8.764 22.002l7.491 7.09"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M39.372 38.5a2 2 0 1 1 0 4h-3.3v-8h3.3a2 2 0 1 1 0 4h0Zm0 0h-3.3"/><circle cx="38.5" cy="38.5" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}