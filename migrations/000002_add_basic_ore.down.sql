DELETE FROM entity
    WHERE code IN
          (
           'iron-ore',
           'copper-ore',
           'coal',
           'gold-ore',
           'silver-ore'
          );