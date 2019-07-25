Pod::Spec.new do |spec|
  spec.name         = 'Gdatx'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://github.com/KunkaYU/go-DATx'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS DATx Client'
  spec.source       = { :git => 'https://github.com/KunkaYU/go-DATx.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/Gdatx.framework'

	spec.prepare_command = <<-CMD
    curl https://gdatxstore.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/Gdatx.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end