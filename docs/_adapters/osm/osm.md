---
layout: default
title: Open Service Mesh
name: Meshery Adapter for Open Service Mesh
mesh_name: Open Service Mesh
version: OSM v0.3.0
port: 10009/tcp
project_status: alpha
github_link: https://github.com/layer5io/meshery-osm
image: /assets/img/service-meshes/osm.svg
permalink: service-meshes/adapters/osm
---

{% include adapter-status.html %}

### Lifecycle management

The {{page.name}} can install **{{page.version}}** of the {{page.mesh_name}} service mesh. Sample applications for {{page.mesh_name}} can also be installed using Meshery. Using the {{page.name}}, you may also run the complete suite of Service Mesh Interface (SMI) conformance tests to validate OSM’s compliance with the SMI specification.

### Install {{ page.mesh_name }}

##### **Choose the Meshery Adapter for {{ page.mesh_name }}**

<a href="{{ site.baseurl }}/assets/img/adapters/osm/osm-adapter.png">
  <img style="width:500px;" src="{{ site.baseurl }}/assets/img/adapters/osm/osm-adapter.png" />
</a>

##### **Click on (+) and choose the `{{page.version}}` of the {{page.mesh_name}} service mesh.**

<a href="{{ site.baseurl }}/assets/img/adapters/osm/osm-install.png">
  <img style="width:500px;" src="{{ site.baseurl }}/assets/img/adapters/osm/osm-install.png" />
</a>


### SMI Conformance Capabiliy

Meshery allows you to analyze the compliance status and functional capabilities of your service mesh. This allows you to compare high-level functional differences between service meshes and verify whether your service mesh is conformant with the SMI specification.

Learn more about the SMI specification and [Meshery's conformance test suite](https://docs.meshery.io/functionality/smi-conformance).

### Sample Applications

The {{ page.name }} does not support Sample Applications yet.

### Suggested Topics

- Examine [Meshery's architecture]({{ site.baseurl }}/architecture) and how adapters fit in as a component.
- Learn more about [Meshery Adapters]({{ site.baseurl }}/architecture/adapters).

