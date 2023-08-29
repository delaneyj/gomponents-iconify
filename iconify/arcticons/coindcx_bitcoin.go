package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoindcxBitcoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.098 32.178c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2s2 .9 2 2v1.3c0 1.1-.9 2-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.295 26.878v5.3m2.638-7v7m-1.101-5.3h2.101" class="f"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.691 31.178c-.3.6-1 1-1.7 1c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2c.7 0 1.4.4 1.7 1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.324 32.178v-3.3c0-1.1-.9-2-2-2s-2 .9-2 2v3.3m0-3.3v-2m-22.647 2c0-1.1.9-2 2-2s2 .9 2 2v1.3c0 1.1-.9 2-2 2s-2-.9-2-2m0 2v-8m21.02 2.7v5.3" class="f"/><circle cx="31.697" cy="24.478" r=".747" fill="currentColor"/><circle cx="16.295" cy="24.478" r=".747" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.708 16.342l-4.081 6.114m-10.814.038V16.38h1c1.694 0 3.08 1.376 3.08 3.057s-1.386 3.057-3.08 3.057h-1Zm9.804-2.063c0 1.147-.924 2.064-2.079 2.064s-2.078-.917-2.078-2.064v-2.063c0-1.146.923-2.064 2.078-2.064s2.002.917 2.002 2.064m1.087-2.026l4.081 6.114m-26.13 0c-.847 0-1.54-.687-1.54-1.528v-.994c0-.84.693-1.528 1.54-1.528s1.54.688 1.54 1.528v.994c0 .84-.693 1.528-1.54 1.528Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.767 18.406v4.05m4.79-.038v-2.523c0-.84-.693-1.528-1.54-1.528s-1.54.688-1.54 1.528v2.523m0-2.523v-1.528" class="f"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.45 20.393c0 1.147-.924 2.064-2.079 2.064s-2.079-.917-2.079-2.064V18.33c0-1.146.924-2.064 2.08-2.064s2 .917 2 2.064"/><circle cx="17.767" cy="16.572" r=".747" fill="currentColor"/>`),
		g.Group(children),
	)
}