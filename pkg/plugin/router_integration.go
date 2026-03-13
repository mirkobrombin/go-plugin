package plugin

import "github.com/mirkobrombin/go-module-router/v2/pkg/registry"

// RegisterRouteProviders inspects registered plugins and forwards any routes they expose.
func RegisterRouteProviders(r *Registry) {
	if r == nil {
		return
	}

	for _, name := range r.Names() {
		p, ok := r.Get(name)
		if !ok {
			continue
		}
		if provider, ok := p.(interface{ Routes() []registry.Route }); ok {
			routes := provider.Routes()
			if len(routes) > 0 {
				registry.RegisterRoutes(func() []registry.Route { return routes })
			}
		}
	}
}
