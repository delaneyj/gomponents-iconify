package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Postnl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.572 5.716c4.444 0 11.584 2.82 16.337 5.442c5.224 2.88 12.584 8.762 12.591 12.818c-.008 4.316-7.534 10.03-12.591 12.82c-4.592 2.53-11.584 5.488-16.175 5.488a6.134 6.134 0 0 1-2.953-.616C7.153 39.667 5.5 31.094 5.5 23.976A44.255 44.255 0 0 1 6.766 13.43c.967-3.76 2.356-6.229 4.015-7.145a5.827 5.827 0 0 1 2.791-.57m1.18 14.43h5.108m-2.554-1.359v-2.935"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.947 17.024c-.367-1.114-1.032-.978-1.436-.584c-.297.29-.102 1.6.495 2.226m3.658-1.642c.367-1.114 1.033-.978 1.436-.584c.297.29.103 1.6-.495 2.226"/><rect width="3.877" height="5.138" x="15.358" y="21.997" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.939"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.185 26.701a2.18 2.18 0 0 0 1.595.434h.435a1.283 1.283 0 0 0 1.281-1.285h0a1.283 1.283 0 0 0-1.281-1.284h-.87a1.283 1.283 0 0 1-1.282-1.285h0a1.283 1.283 0 0 1 1.282-1.284h.435a2.18 2.18 0 0 1 1.595.433m2.887-2.032v5.767a.97.97 0 0 0 .97.97h.29m-2.278-5.138h2.036m6.024 5.138v-3.2a1.939 1.939 0 0 0-1.939-1.938h0a1.939 1.939 0 0 0-1.939 1.939v3.199m0-3.199v-1.939M9.6 25.168a1.939 1.939 0 0 0 1.939 1.939h0a1.939 1.939 0 0 0 1.939-1.94v-1.26a1.939 1.939 0 0 0-1.94-1.938h0A1.939 1.939 0 0 0 9.6 23.908m0-1.939v7.755M36.182 19.38v7.755"/><circle cx="17.306" cy="13.972" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}