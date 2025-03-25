"""
Captura de dados de temperatura, carga de CPU e valores de fan para induzir uma fórmula 
de predição de fan a partir de temperatura e carga de CPU.

Regressão linear
"""
import pandas as pd
from sklearn.linear_model import LinearRegression
from sklearn.model_selection import train_test_split
from sklearn.metrics import mean_squared_error

# 1. Preparação dos Dados
df = pd.read_csv('seu_arquivo.csv')
X = df[['cpu_load', 'temperature']]
y = df['fan']

# 2. Divisão dos dados em treino e teste
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

# 3. Modelagem
model = LinearRegression()
model.fit(X_train, y_train)

# 4. Avaliação
y_pred = model.predict(X_test)
mse = mean_squared_error(y_test, y_pred)
print(f'MSE: {mse}')

# 5. Fórmula
b0 = model.intercept_
b1 = model.coef_[0]
b2 = model.coef_[1]
print(f'Fórmula: fan = {b0:.2f} + {b1:.2f} * cpu_load + {b2:.2f} * temperature')