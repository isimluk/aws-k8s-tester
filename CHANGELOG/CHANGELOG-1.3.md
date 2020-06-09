

<hr>


## [v1.3.1](https://github.com/aws/aws-k8s-tester/releases/tag/v1.3.1) (2020-06)

See [code changes](https://github.com/aws/aws-k8s-tester/compare/v1.3.0...v1.3.1).

TODO

### Dependency

- Upgrade [`github.com/aws/aws-sdk-go`](https://github.com/aws/aws-sdk-go/releases) from [`v1.31.10`](https://github.com/aws/aws-sdk-go/releases/tag/v1.31.10) to [`v1.31.12`](https://github.com/aws/aws-sdk-go/releases/tag/v1.31.12).
- Upgrade [`helm.sh/helm/v3`](https://github.com/helm/helm/releases) from [`v3.2.1`](https://github.com/helm/helm/releases/tag/v3.2.1) to [`v3.2.2`](https://github.com/helm/helm/releases/tag/v3.2.2).

### Go

- Compile with [*Go 1.14.4*](https://golang.org/doc/devel/release.html#go1.14).


<hr>


## [v1.3.0](https://github.com/aws/aws-k8s-tester/releases/tag/v1.3.0) (2020-06-03)

See [code changes](https://github.com/aws/aws-k8s-tester/compare/v1.2.9...v1.3.0).

### `eksconfig`

- Add [`AddOnCUDAVectorAdd`](https://github.com/aws/aws-k8s-tester/pull/89).
  - Can be enabled via `AWS_K8S_TESTER_EKS_ADD_ON_CUDA_VECTOR_ADD_ENABLE`.

### `eks`

- Add [`eks/cuda-vector-add`](https://github.com/aws/aws-k8s-tester/pull/89).
- Improve [`eks/cuda-vector-add` output checks](https://github.com/aws/aws-k8s-tester/commit/75ca40a81845eba3a3b2246fb7a67f0dcc82bf8b).

### Dependency

- Upgrade [`e2e/tester/pkg` `kops` dependency to `1.17.6`](https://github.com/aws/aws-k8s-tester/pull/88).
- Upgrade [`github.com/aws/aws-sdk-go`](https://github.com/aws/aws-sdk-go/releases) from [`v1.31.9`](https://github.com/aws/aws-sdk-go/releases/tag/v1.31.9) to [`v1.31.10`](https://github.com/aws/aws-sdk-go/releases/tag/v1.31.10).

### Go

- Compile with [*Go 1.14.4*](https://golang.org/doc/devel/release.html#go1.14).


<hr>
