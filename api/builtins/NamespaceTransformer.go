// Code generated by pluginator on NamespaceTransformer; DO NOT EDIT.
// pluginator {unknown  1970-01-01T00:00:00Z  }

package builtins

import (
	"fmt"

	"sigs.k8s.io/kustomize/api/filters/namespace"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/yaml"
)

// Change or set the namespace of non-cluster level resources.
type NamespaceTransformerPlugin struct {
	types.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	FieldSpecs       []types.FieldSpec `json:"fieldSpecs,omitempty" yaml:"fieldSpecs,omitempty"`
}

func (p *NamespaceTransformerPlugin) Config(
	_ *resmap.PluginHelpers, c []byte) (err error) {
	p.Namespace = ""
	p.FieldSpecs = nil
	return yaml.Unmarshal(c, p)
}

func (p *NamespaceTransformerPlugin) Transform(m resmap.ResMap) error {
	if len(p.Namespace) == 0 {
		return nil
	}
	for _, r := range m.Resources() {
		if r.IsEmpty() {
			// Don't mutate empty objects?
			continue
		}
		err := r.ApplyFilter(namespace.Filter{
			Namespace: p.Namespace,
			FsSlice:   p.FieldSpecs,
		})
		if err != nil {
			return err
		}
		matches := m.GetMatchingResourcesByCurrentId(r.CurId().Equals)
		if len(matches) != 1 {
			return fmt.Errorf(
				"namespace transformation produces ID conflict: %+v", matches)
		}
	}
	return nil
}

func NewNamespaceTransformerPlugin() resmap.TransformerPlugin {
	return &NamespaceTransformerPlugin{}
}
