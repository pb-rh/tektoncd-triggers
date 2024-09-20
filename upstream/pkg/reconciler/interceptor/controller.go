/*
Copyright 2021 The Tekton Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package interceptor

import (
	"context"

	interceptorinformer "github.com/tektoncd/triggers/pkg/client/injection/informers/triggers/v1alpha1/interceptor"
	interceptorreconciler "github.com/tektoncd/triggers/pkg/client/injection/reconciler/triggers/v1alpha1/interceptor"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
)

func NewController() func(context.Context, configmap.Watcher) *controller.Impl {
	return func(ctx context.Context, _ configmap.Watcher) *controller.Impl {
		interceptorInformer := interceptorinformer.Get(ctx)
		reconciler := &Reconciler{}

		impl := interceptorreconciler.NewImpl(ctx, reconciler, func(_ *controller.Impl) controller.Options {
			return controller.Options{
				AgentName: ControllerName,
			}
		})

		if _, err := interceptorInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue)); err != nil {
			logging.FromContext(ctx).Panicf("Couldn't register Interceptor informer event handler: %w", err)
		}

		return impl
	}
}
